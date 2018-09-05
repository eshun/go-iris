package models

type AdPosition struct {
	Id     int    `json:"id" xorm:"not null pk autoincr TINYINT(3)"`
	Name   string `json:"name" xorm:"not null default '' VARCHAR(60)"`
	Width  int    `json:"width" xorm:"not null default 0 SMALLINT(5)"`
	Height int    `json:"height" xorm:"not null default 0 SMALLINT(5)"`
	Desc   string `json:"desc" xorm:"not null default '' VARCHAR(255)"`
}
