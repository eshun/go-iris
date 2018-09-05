package models

type Product struct {
	Id                    int    `json:"id" xorm:"not null pk autoincr MEDIUMINT(8)"`
	GoodsId               int    `json:"goods_id" xorm:"not null default 0 MEDIUMINT(8)"`
	GoodsSpecificationIds string `json:"goods_specification_ids" xorm:"not null default '' VARCHAR(50)"`
	GoodsSn               string `json:"goods_sn" xorm:"not null default '' VARCHAR(60)"`
	GoodsNumber           int    `json:"goods_number" xorm:"not null default 0 MEDIUMINT(8)"`
	RetailPrice           string `json:"retail_price" xorm:"not null default 0.00 DECIMAL(10,2)"`
}
