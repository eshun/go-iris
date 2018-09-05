package models

type Ad struct {
	Id           int    `json:"id" xorm:"not null pk autoincr SMALLINT(5)"`
	AdPositionId int    `json:"ad_position_id" xorm:"not null default 0 index SMALLINT(5)"`
	MediaType    int    `json:"media_type" xorm:"not null default 0 TINYINT(3)"`
	Name         string `json:"name" xorm:"not null default '' VARCHAR(60)"`
	Link         string `json:"link" xorm:"not null default '' VARCHAR(255)"`
	ImageUrl     string `json:"image_url" xorm:"not null TEXT"`
	Content      string `json:"content" xorm:"not null default '' VARCHAR(255)"`
	EndTime      int    `json:"end_time" xorm:"not null default 0 INT(11)"`
	Enabled      int    `json:"enabled" xorm:"not null default 1 index TINYINT(3)"`
}
