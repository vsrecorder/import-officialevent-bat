package models

type EventDetailSearch struct {
	Code        uint16      `json:"code"`
	EventDetail EventDetail `json:"event"`
}
