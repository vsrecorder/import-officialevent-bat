package models

type EventDetail struct {
	Id              uint            `json:"id"`
	Title           string          `json:"eventTitle"`
	Address         string          `json:"address"`
	Venue           string          `json:"venue"`
	Date            EventDetailDate `json:"eventDate"`
	StartedAt       EventDetailDate `json:"eventStartedAt"`
	EndedAt         EventDetailDate `json:"eventEndedAt"`
	DeckCount       string          `json:"deckCount"`
	TypeId          uint            `json:"eventTypeId"`
	TypeName        string          `json:"eventTypeName"`
	CSPFlg          uint8           `json:"cspFlg"`
	LeagueId        uint            `json:"eventLeagueId"`
	LeagueTitle     string          `json:"leagueTitle"`
	RegulationId    uint            `json:"regulationId"`
	RegulationTitle string          `json:"regulationTitle"`
	Capacity        uint            `json:"capacity"`
	AttrId          uint            `json:"eventAttrId"`
	ShopId          uint            `json:"shopId"`
	ShopName        string          `json:"shopName"`
}
