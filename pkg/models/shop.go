package models

type Shop struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Term          uint   `json:"term"`
	ZipCode       string `json:"zip_code"`
	PrefectureId  uint8  `json:"prefecture_id"`
	Address       string `json:"address"`
	Tel           string `json:"tel"`
	Access        string `json:"access"`
	BusinessHours string `json:"business_hours"`
	Url           string `json:"url"`
	GeoCoding     string `json:"geo_coding"`
}
