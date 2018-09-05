package models

type Order struct {
	Id             int    `json:"id" xorm:"not null pk autoincr MEDIUMINT(8)"`
	OrderSn        string `json:"order_sn" xorm:"not null default '' unique VARCHAR(20)"`
	UserId         int    `json:"user_id" xorm:"not null default 0 index MEDIUMINT(8)"`
	OrderStatus    int    `json:"order_status" xorm:"not null default 0 index TINYINT(1)"`
	ShippingStatus int    `json:"shipping_status" xorm:"not null default 0 index TINYINT(1)"`
	PayStatus      int    `json:"pay_status" xorm:"not null default 0 index TINYINT(1)"`
	Consignee      string `json:"consignee" xorm:"not null default '' VARCHAR(60)"`
	Country        int    `json:"country" xorm:"not null default 0 SMALLINT(5)"`
	Province       int    `json:"province" xorm:"not null default 0 SMALLINT(5)"`
	City           int    `json:"city" xorm:"not null default 0 SMALLINT(5)"`
	District       int    `json:"district" xorm:"not null default 0 SMALLINT(5)"`
	Address        string `json:"address" xorm:"not null default '' VARCHAR(255)"`
	Mobile         string `json:"mobile" xorm:"not null default '' VARCHAR(60)"`
	Postscript     string `json:"postscript" xorm:"not null default '' VARCHAR(255)"`
	ShippingId     int    `json:"shipping_id" xorm:"not null default 0 index TINYINT(3)"`
	ShippingName   string `json:"shipping_name" xorm:"not null default '' VARCHAR(120)"`
	PayId          int    `json:"pay_id" xorm:"not null default 0 index TINYINT(3)"`
	PayName        string `json:"pay_name" xorm:"not null default '' VARCHAR(120)"`
	ShippingFee    string `json:"shipping_fee" xorm:"not null default 0.00 DECIMAL(10,2)"`
	ActualPrice    string `json:"actual_price" xorm:"not null default 0.00 DECIMAL(10,2)"`
	Integral       int    `json:"integral" xorm:"not null default 0 INT(10)"`
	IntegralMoney  string `json:"integral_money" xorm:"not null default 0.00 DECIMAL(10,2)"`
	OrderPrice     string `json:"order_price" xorm:"not null default 0.00 DECIMAL(10,2)"`
	GoodsPrice     string `json:"goods_price" xorm:"not null default 0.00 DECIMAL(10,2)"`
	AddTime        int    `json:"add_time" xorm:"not null default 0 INT(11)"`
	ConfirmTime    int    `json:"confirm_time" xorm:"not null default 0 INT(11)"`
	PayTime        int    `json:"pay_time" xorm:"not null default 0 INT(11)"`
	FreightPrice   int    `json:"freight_price" xorm:"not null default 0 INT(10)"`
	CouponId       int    `json:"coupon_id" xorm:"not null default 0 MEDIUMINT(8)"`
	ParentId       int    `json:"parent_id" xorm:"not null default 0 MEDIUMINT(8)"`
	CouponPrice    string `json:"coupon_price" xorm:"not null DECIMAL(10,2)"`
	CallbackStatus string `json:"callback_status" xorm:"default 'true' ENUM('true','false')"`
}
