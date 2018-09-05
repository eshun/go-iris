package models

type Channel struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name      string `json:"name" xorm:"not null default '' VARCHAR(45)"`
	Url       string `json:"url" xorm:"not null default '' VARCHAR(255)"`
	IconUrl   string `json:"icon_url" xorm:"not null default '' VARCHAR(255)"`
	SortOrder int    `json:"sort_order" xorm:"not null default 10 INT(4)"`
}
