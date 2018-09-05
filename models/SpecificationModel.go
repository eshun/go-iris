package models

type Specification struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name      string `json:"name" xorm:"not null default '' VARCHAR(60)"`
	SortOrder int    `json:"sort_order" xorm:"not null default 0 TINYINT(3)"`
}
