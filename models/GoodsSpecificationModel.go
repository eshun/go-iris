package models

type GoodsSpecification struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	GoodsId         int    `json:"goods_id" xorm:"not null default 0 index INT(11)"`
	SpecificationId int    `json:"specification_id" xorm:"not null default 0 index INT(11)"`
	Value           string `json:"value" xorm:"not null default '' VARCHAR(50)"`
	PicUrl          string `json:"pic_url" xorm:"not null default '' VARCHAR(255)"`

	Name     string `json:"name" xorm:"-"`
}
