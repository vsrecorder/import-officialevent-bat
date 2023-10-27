package models

// 公式サイトから取得できるショップ情報のshopIdとtermが数値だったり文字列だったりしているため、
// 取得しないようにダミーのタグをつけた
// https://players.pokemon-card.com/shop?shop_id=%d&targetMonth=%s

type Shop struct {
	Id             uint   `json:"dumy_shopId"`
	Name           string `json:"shopName"`
	Term           uint   `json:"dumy_term"`
	ZipCode        string `json:"zip"`
	PrefectureName string `json:"pref"`
	Address        string `json:"addr"`
	Tel            string `json:"tel"`
	Access         string `json:"access"`
	BusinessHours  string `json:"businessHours"`
	Url            string `json:"url"`
	GeoCoding      string `json:"geocoding"`
}
