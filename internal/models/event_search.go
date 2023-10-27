package models

type EventSearch struct {
	Code       uint16  `json:"code"`
	Event      []Event `json:"event"`
	EventCount uint16  `json:"eventCount"`
}
