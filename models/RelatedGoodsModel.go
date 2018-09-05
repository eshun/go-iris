package models

type RelatedGoods struct {
	Id             int `json:"id" xorm:"not null pk autoincr INT(11)"`
	GoodsId        int `json:"goods_id" xorm:"not null default 0 INT(11)"`
	RelatedGoodsId int `json:"related_goods_id" xorm:"not null default 0 INT(11)"`
}
