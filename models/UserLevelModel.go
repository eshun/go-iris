package models

type UserLevel struct {
	Id          int    `json:"id" xorm:"not null pk autoincr TINYINT(3)"`
	Name        string `json:"name" xorm:"not null default '' VARCHAR(30)"`
	Description string `json:"description" xorm:"not null default '' VARCHAR(255)"`
}
