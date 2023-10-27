package models

type Event struct {
	DateId         uint   `json:"date_id"`
	ShopId         uint   `json:"shop_id"`
	EventDateParms string `json:"event_date_params"`
	EventStartedAt string `json:"event_started_at"`
	EventEndedAt   string `json:"event_ended_at"`
	PrefectureName string `json:"prefecture_name"`
	DeckCount      string `json:"deck_count"`
	Address        string `json:"address"`
	EventTitle     string `json:"event_title"`
	EventHoldingId uint   `json:"event_holding_id"`
	EventType      uint   `json:"event_type"`
	CSPFlg         uint8  `json:"csp_flg"`
	EventLeague    uint   `json:"event_league"`
	Regulation     string `json:"regulation"`
	EntryFee       string `json:"entry_fee"`
	Capacity       uint   `json:"capacity"`
	ShopName       string `json:"shop_name"`
	ShopTerm       uint   `json:"shop_term"`
	LeagueName     string `json:"leagueName"`
	EventAttrId    uint   `json:"event_attr_id"`
	TrainersFlg    uint8  `json:"trainers_flg"`
	HolidayFlg     uint8  `json:"holiday_flg"`
}
