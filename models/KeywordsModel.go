package models

type Keywords struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Keyword   string `json:"keyword" xorm:"not null pk default '' VARCHAR(90)"`
	IsHot     int    `json:"is_hot" xorm:"not null default 0 TINYINT(1)"`
	IsDefault int    `json:"is_default" xorm:"not null default 0 TINYINT(1)"`
	IsShow    int    `json:"is_show" xorm:"not null default 1 TINYINT(1)"`
	SortOrder int    `json:"sort_order" xorm:"not null default 100 INT(11)"`
	SchemeUrl string `json:"scheme_url" xorm:"not null default '' VARCHAR(255)"`
	Type      int    `json:"type" xorm:"not null default 0 INT(11)"`
}
