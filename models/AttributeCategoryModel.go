package models

type AttributeCategory struct {
	Id      int    `json:"id" xorm:"not null pk INT(11)"`
	Name    string `json:"name" xorm:"not null default '' VARCHAR(60)"`
	Enabled int    `json:"enabled" xorm:"not null default 1 TINYINT(1)"`
}
