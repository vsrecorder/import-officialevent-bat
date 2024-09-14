package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/vsrecorder/import-officialevent-bat/infrastructures"
	"github.com/vsrecorder/import-officialevent-bat/internal/models"
	daos "github.com/vsrecorder/import-officialevent-bat/pkg/models"
	"gorm.io/gorm"
)

// 公式イベントの数を取得
func getEventCount(t time.Time) (uint16, error) {
	year := uint16(t.Year())
	month := uint8(t.Month())
	day := uint8(t.Day())

	var eventCount uint16

	res, err := http.Get(fmt.Sprintf("https://players.pokemon-card.com/event_search/count?start_date=%d/%02d/%02d&end_date=%d/%02d/%02d", year, month, day, year, month, day))
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	var eventCountSearch models.EventCountSearch
	if err := json.Unmarshal(body, &eventCountSearch); err != nil {
		return 0, err
	}

	eventCount = uint16(eventCountSearch.Count)

	return eventCount, err
}

// 公式イベントデータを取得
func getEvent(t time.Time) ([]models.Event, error) {
	year := uint16(t.Year())
	month := uint8(t.Month())
	day := uint8(t.Day())
	eventCount, err := getEventCount(t)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(fmt.Sprintf("https://players.pokemon-card.com/event_search?start_date=%d/%02d/%02d&end_date=%d/%02d/%02d&limit=%d", year, month, day, year, month, day, eventCount))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var eventSearch models.EventSearch
	if err := json.Unmarshal(body, &eventSearch); err != nil {
		return nil, err
	}

	return eventSearch.Event, nil
}

func stringToTime(str string) time.Time {
	var layout = "2006-01-02 15:04:00.000000"

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	t, err := time.ParseInLocation(layout, str, jst)
	if err != nil {
		panic(err)
	}

	return t
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	userName := os.Getenv("DB_USER_NAME")
	password := os.Getenv("DB_PASSWORD")
	dbHostname := os.Getenv("DB_HOSTNAME")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := infrastructures.NewMySQL(userName, password, dbHostname, dbPort, dbName)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 公式イベントデータをWebから取得
	var events []models.Event
	events, err = getEvent(time.Now())
	if err != nil {
		panic(err)
	}

	for _, event := range events {
		// ショップが既にDBに存在するか確認
		// ショップがDBに存在しない場合、DBに登録する
		if result := db.Where(&models.Shop{Id: event.ShopId}).First(&models.Shop{}); errors.Is(result.Error, gorm.ErrRecordNotFound) {

			// ショップ情報の取得
			// ↓の返り値のshopIdとtermが数値だったり、文字列だったりしているから気をつける
			res, err := http.Get(fmt.Sprintf("https://players.pokemon-card.com/shop?shop_id=%d&targetMonth=%s", event.ShopId, event.EventDateParms))
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			var shopSearch models.ShopSearch
			if err := json.Unmarshal(body, &shopSearch); err != nil {
				panic(err)
			}

			var pref models.Prefectures
			db.Where("name = ?", shopSearch.Shop.PrefectureName).First(&pref)

			{
				var shop daos.Shop

				shop.Id = event.ShopId
				shop.Name = shopSearch.Shop.Name
				shop.ZipCode = shopSearch.Shop.ZipCode
				shop.PrefectureId = pref.Id
				shop.Address = shopSearch.Shop.Address
				shop.Tel = shopSearch.Shop.Tel
				shop.Access = shopSearch.Shop.Access
				shop.BusinessHours = shopSearch.Shop.BusinessHours
				shop.Url = shopSearch.Shop.Url
				shop.GeoCoding = shopSearch.Shop.GeoCoding

				// 返り値のshopIdとtermが数値だったり、文字列だったりしているから処理する
				var shopTermStringSearch models.ShopTermStringSearch
				if err := json.Unmarshal(body, &shopTermStringSearch); err != nil {
					var shopTermUintSearch models.ShopTermUintSearch
					if err := json.Unmarshal(body, &shopTermUintSearch); err != nil {
						panic(err)
					} else {
						shop.Term = shopTermUintSearch.ShopTermUint.Term
					}
				} else {
					i, _ := strconv.Atoi(shopTermStringSearch.ShopTermString.Term)
					shop.Term = uint(i)
				}

				db.Create(&shop)
			}
		}

		// 公式イベントをDBに登録する
		{
			// 公式イベントの詳細情報の取得
			res, err := http.Get(fmt.Sprintf("https://players.pokemon-card.com/event_detail_search?event_holding_id=%d", event.EventHoldingId))
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			var eventDetailSearch models.EventDetailSearch
			if err := json.Unmarshal(body, &eventDetailSearch); err != nil {
				panic(err)
			}
			eventDetail := eventDetailSearch.EventDetail
			fmt.Println("event_holding_id:", eventDetail.Id)

			var officialEvent daos.OfficialEvent
			officialEvent.Id = eventDetail.Id

			// オーガナイザーイベントの場合
			if eventDetail.TypeId == 6 {
				officialEvent.Title = eventDetail.OrgTitle
			} else {
				officialEvent.Title = eventDetail.Title
			}

			officialEvent.Address = eventDetail.Address
			officialEvent.Venue = eventDetail.Venue

			jst, err := time.LoadLocation("Asia/Tokyo")
			if err != nil {
				panic(err)
			}

			d := stringToTime(eventDetail.Date.Date)
			officialEvent.Date = d

			if eventDetail.StartedAt.Date != "" {
				s := stringToTime(eventDetail.StartedAt.Date)
				startedDate := time.Date(d.Year(), d.Month(), d.Day(), s.Hour(), s.Minute(), s.Second(), 0, jst)
				officialEvent.StartedAt = startedDate
			} else {
				officialEvent.StartedAt = d
			}

			if eventDetail.EndedAt.Date != "" {
				e := stringToTime(eventDetail.EndedAt.Date)
				endedDate := time.Date(d.Year(), d.Month(), d.Day(), e.Hour(), e.Minute(), e.Second(), 0, jst)
				officialEvent.EndedAt = endedDate
			} else {
				officialEvent.EndedAt = d
			}

			officialEvent.DeckCount = eventDetail.DeckCount
			officialEvent.TypeId = eventDetail.TypeId
			officialEvent.TypeName = eventDetail.TypeName

			if eventDetail.CSPFlg == 1 {
				officialEvent.CSPFlg = true
			} else {
				officialEvent.CSPFlg = false
			}

			officialEvent.LeagueId = eventDetail.LeagueId
			officialEvent.LeagueTitle = eventDetail.LeagueTitle
			officialEvent.RegulationId = eventDetail.RegulationId
			officialEvent.RegulationTitle = eventDetail.RegulationTitle
			officialEvent.Capacity = eventDetail.Capacity
			officialEvent.AttrId = eventDetail.AttrId
			officialEvent.ShopId = eventDetail.ShopId
			officialEvent.ShopName = eventDetail.ShopName

			if result := db.Where(&models.OfficialEvent{Id: officialEvent.Id}).First(&models.OfficialEvent{}); errors.Is(result.Error, gorm.ErrRecordNotFound) {
				db.Create(&officialEvent)
			}
		}
	}

}
