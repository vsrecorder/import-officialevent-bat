package models

import (
	"time"
)

type OfficialEvent struct {
	Id          uint
	Name        string
	Address     string
	Date        time.Time
	StartedAt   time.Time
	EndedAt     time.Time
	DeckCount   string
	Type        uint8
	CSPFlg      uint8
	League      uint16
	LeagueName  string
	Regulation  string
	EntryFee    string
	Capacity    uint16
	AttrId      uint8
	TrainersFlg uint8
	HolidayFlg  uint8
	DateId      uint
	ShopId      uint
}
