package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/vsrecorder/import-officialevent-bat/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 公式イベントの数を取得
func getEventCount(t time.Time) (uint16, error) {
	year := uint16(t.Year())
	month := uint8(t.Month())
	day := uint8(t.Day())

	var eventCount uint16

	res, err := http.Get(fmt.Sprintf("https://players.pokemon-card.com/event_search/count?start_date=%d/%d/%d&end_date=%d/%d/%d", year, month, day, year, month, day))
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	var eventCountSearch model.EventCountSearch
	if err := json.Unmarshal(body, &eventCountSearch); err != nil {
		return 0, err
	}

	eventCount = uint16(eventCountSearch.Count)

	return eventCount, err
}

// 公式イベントデータを取得
func getEvent(t time.Time) ([]model.Event, error) {
	year := uint16(t.Year())
	month := uint8(t.Month())
	day := uint8(t.Day())
	eventCount, err := getEventCount(t)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(fmt.Sprintf("https://players.pokemon-card.com/event_search?start_date=%d/%d/%d&end_date=%d/%d/%d&limit=%d", year, month, day, year, month, day, eventCount))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var eventSearch model.EventSearch
	if err := json.Unmarshal(body, &eventSearch); err != nil {
		return nil, err
	}

	return eventSearch.Event, nil
}

func stringToTime(str string) time.Time {
	var layout = "20060102 15:04"

	jst, _ := time.LoadLocation("Asia/Tokyo")
	t, _ := time.ParseInLocation(layout, str, jst)
	return t
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	userName := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB_NAME")
	mysqlHostname := os.Getenv("MYSQL_HOSTNAME")

	// DB接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, mysqlHostname, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 公式イベントデータを読み込む
	var events []model.Event

	// ファイルから読み込む
	if len(os.Args) == 2 {

		fileName := os.Args[1]
		fdata, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		var eventSearch model.EventSearch
		if err := json.Unmarshal([]byte(fdata), &eventSearch); err != nil {
			panic(err)
		}
		events = eventSearch.Event
	} else {
		// Webから読み込む
		// 公式イベントデータを取得
		events, err = getEvent(time.Now())
		if err != nil {
			panic(err)
		}
	}

	for _, event := range events {
		// ショップが既にDBに存在するか確認
		// ショップがDBに存在しない場合、DBに登録する
		if result := db.Where(&model.Shop{Id: event.ShopId}).First(&model.Shop{}); errors.Is(result.Error, gorm.ErrRecordNotFound) {

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

			var shopSearch model.ShopSearch
			if err := json.Unmarshal(body, &shopSearch); err != nil {
				panic(err)
			}

			var pref model.Prefectures
			db.Where("name = ?", event.PrefectureName).First(&pref)

			{
				var newShop model.Shop
				newShop.Id = event.ShopId
				newShop.Name = shopSearch.Shop.Name
				newShop.Term = event.ShopTerm
				newShop.ZipCode = shopSearch.Shop.ZipCode
				newShop.Address = shopSearch.Shop.Address
				newShop.PrefectureId = pref.Id
				newShop.Address = shopSearch.Shop.Address
				newShop.Tel = shopSearch.Shop.Tel
				newShop.Access = shopSearch.Shop.Access
				newShop.BusinessHours = shopSearch.Shop.BusinessHours
				newShop.Url = shopSearch.Shop.Url
				newShop.GeoCoding = shopSearch.Shop.GeoCoding

				db.Create(&newShop)
			}
		}

		// 公式イベントをDBに登録する
		{
			var officialEvent model.OfficialEvent
			officialEvent.Id = event.EventHoldingId
			officialEvent.Name = event.EventTitle
			officialEvent.Address = event.Address
			officialEvent.DeckCount = event.DeckCount
			officialEvent.Type = event.EventType
			officialEvent.CSPFlg = event.CSPFlg
			officialEvent.League = event.EventLeague
			officialEvent.LeagueName = event.LeagueName
			officialEvent.Regulation = event.Regulation
			officialEvent.EntryFee = event.EntryFee
			officialEvent.Capacity = event.Capacity
			officialEvent.AttrId = event.EventAttrId
			officialEvent.TrainersFlg = event.TrainersFlg
			officialEvent.HolidayFlg = event.HolidayFlg
			officialEvent.DateId = event.DateId
			officialEvent.ShopId = event.ShopId
			dateTime := stringToTime(event.EventDateParms + " 00:00")
			officialEvent.Date = dateTime
			if event.EventStartedAt != "" {
				startedAtTime := stringToTime(event.EventDateParms + " " + event.EventStartedAt)
				officialEvent.StartedAt = startedAtTime
			} else {
				officialEvent.StartedAt = dateTime
			}
			if event.EventEndedAt != "" {
				endedAtTime := stringToTime(event.EventDateParms + " " + event.EventEndedAt)
				officialEvent.EndedAt = endedAtTime
			} else {
				officialEvent.EndedAt = dateTime

			}

			if result := db.Where(&model.OfficialEvent{Id: officialEvent.Id}).First(&model.OfficialEvent{}); errors.Is(result.Error, gorm.ErrRecordNotFound) {
				db.Create(&officialEvent)
			}
		}
	}
}
