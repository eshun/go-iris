package models

type GoodsAttribute struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	GoodsId     int    `json:"goods_id" xorm:"not null default 0 index INT(11)"`
	AttributeId int    `json:"attribute_id" xorm:"not null default 0 index INT(11)"`
	Value       string `json:"value" xorm:"not null TEXT"`

	Name string `json:"name" xorm:"-"`
}
