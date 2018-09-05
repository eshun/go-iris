package models

type GoodsGallery struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	GoodsId   int    `json:"goods_id" xorm:"not null default 0 index INT(11)"`
	ImgUrl    string `json:"img_url" xorm:"not null default '' VARCHAR(255)"`
	ImgDesc   string `json:"img_desc" xorm:"not null default '' VARCHAR(255)"`
	SortOrder int    `json:"sort_order" xorm:"not null default 5 INT(11)"`
}
