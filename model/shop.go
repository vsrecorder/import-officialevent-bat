package model

type Shop struct {
	Id            uint
	Name          string `json:"shopName"`
	Term          uint8  `json:"dumy_term"`
	ZipCode       string `json:"zip"`
	PrefectureId  uint8
	Address       string `json:"addr"`
	Tel           string `json:"tel"`
	Access        string `json:"access"`
	BusinessHours string `json:"businessHours"`
	Url           string `json:"url"`
	GeoCoding     string `json:"geocoding"`
}
