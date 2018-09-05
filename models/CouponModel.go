package models

type Coupon struct {
	Id             int    `json:"id" xorm:"not null pk autoincr SMALLINT(5)"`
	Name           string `json:"name" xorm:"not null default '' VARCHAR(60)"`
	TypeMoney      string `json:"type_money" xorm:"not null default 0.00 DECIMAL(10,2)"`
	SendType       int    `json:"send_type" xorm:"not null default 0 TINYINT(3)"`
	MinAmount      string `json:"min_amount" xorm:"not null default 0.00 DECIMAL(10,2)"`
	MaxAmount      string `json:"max_amount" xorm:"not null default 0.00 DECIMAL(10,2)"`
	SendStartDate  int    `json:"send_start_date" xorm:"not null default 0 INT(11)"`
	SendEndDate    int    `json:"send_end_date" xorm:"not null default 0 INT(11)"`
	UseStartDate   int    `json:"use_start_date" xorm:"not null default 0 INT(11)"`
	UseEndDate     int    `json:"use_end_date" xorm:"not null default 0 INT(11)"`
	MinGoodsAmount string `json:"min_goods_amount" xorm:"not null default 0.00 DECIMAL(10,2)"`
}
