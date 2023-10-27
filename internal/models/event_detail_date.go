package models

type EventDetailDate struct {
	Date         string `json:"date"`
	TimeZornType uint   `json:"timezone_type"`
	TimeZone     string `json:"timezone"`
}
