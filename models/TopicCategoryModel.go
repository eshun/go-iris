package models

type TopicCategory struct {
	Id     int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Title  string `json:"title" xorm:"not null default '' VARCHAR(255)"`
	PicUrl string `json:"pic_url" xorm:"not null default '' VARCHAR(255)"`
}
