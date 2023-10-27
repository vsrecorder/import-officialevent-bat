package models

// 公式サイトから取得できるショップ情報のtermが数値だったり文字列だったりしているため、
// stringに決め打ち
// https://players.pokemon-card.com/shop?shop_id=%d&targetMonth=%s

type ShopTermString struct {
	Id             uint   `json:"dumy_shopId"`
	Name           string `json:"shopName"`
	Term           string `json:"term"`
	ZipCode        string `json:"zip"`
	PrefectureName string `json:"pref"`
	Address        string `json:"addr"`
	Tel            string `json:"tel"`
	Access         string `json:"access"`
	BusinessHours  string `json:"businessHours"`
	Url            string `json:"url"`
	GeoCoding      string `json:"geocoding"`
}
