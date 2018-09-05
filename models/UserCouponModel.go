package models

type UserCoupon struct {
	Id           int    `json:"id" xorm:"not null pk autoincr MEDIUMINT(8)"`
	CouponId     int    `json:"coupon_id" xorm:"not null default 0 TINYINT(3)"`
	CouponNumber string `json:"coupon_number" xorm:"not null default '' VARCHAR(20)"`
	UserId       int    `json:"user_id" xorm:"not null default 0 index INT(11)"`
	UsedTime     int    `json:"used_time" xorm:"not null default 0 INT(10)"`
	OrderId      int    `json:"order_id" xorm:"not null default 0 MEDIUMINT(8)"`
}
