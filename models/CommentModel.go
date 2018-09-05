package models

type Comment struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	TypeId     int    `json:"type_id" xorm:"not null default 0 TINYINT(3)"`
	ValueId    int    `json:"value_id" xorm:"not null default 0 index INT(11)"`
	Content    string `json:"content" xorm:"not null VARCHAR(6550)"`
	AddTime    int64  `json:"add_time" xorm:"not null default 0 BIGINT(12)"`
	Status     int    `json:"status" xorm:"not null default 0 TINYINT(3)"`
	UserId     int    `json:"user_id" xorm:"not null default 0 INT(11)"`
	NewContent string `json:"new_content" xorm:"not null default '' VARCHAR(6550)"`
}
