package models

type SearchHistory struct {
	Id      int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Keyword string `json:"keyword" xorm:"not null CHAR(50)"`
	From    string `json:"from" xorm:"not null default '' VARCHAR(45)"`
	AddTime int    `json:"add_time" xorm:"not null default 0 INT(11)"`
	UserId  string `json:"user_id" xorm:"VARCHAR(45)"`
}
