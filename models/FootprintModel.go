package models

type Footprint struct {
	Id      int `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserId  int `json:"user_id" xorm:"not null default 0 INT(11)"`
	GoodsId int `json:"goods_id" xorm:"not null default 0 INT(11)"`
	AddTime int `json:"add_time" xorm:"not null default 0 INT(11)"`
}
