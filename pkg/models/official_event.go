package models

import "time"

type OfficialEvent struct {
	Id              uint      `json:"id"`
	Title           string    `json:"title"`
	Address         string    `json:"address"`
	Venue           string    `json:"venue"`
	Date            time.Time `json:"date"`
	StartedAt       time.Time `json:"started_at"`
	EndedAt         time.Time `json:"ended_at"`
	DeckCount       string    `json:"deck_count"`
	TypeId          uint      `json:"type_dd"`
	TypeName        string    `json:"type_name"`
	CSPFlg          bool      `json:"csp_flg"`
	LeagueId        uint      `json:"league_id"`
	LeagueTitle     string    `json:"league_title"`
	RegulationId    uint      `json:"regulation_id"`
	RegulationTitle string    `json:"regulation_title"`
	Capacity        uint      `json:"capacity"`
	AttrId          uint      `json:"attr_id"`
	ShopId          uint      `json:"shop_id"`
	ShopName        string    `json:"shop_name"`
}
