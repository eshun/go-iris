package models

type Topic struct {
	Id              int    `json:"id" xorm:"not null autoincr index INT(10)"`
	Title           string `json:"title" xorm:"not null default '' VARCHAR(255)"`
	Content         string `json:"content" xorm:"TEXT"`
	Avatar          string `json:"avatar" xorm:"not null default '' VARCHAR(255)"`
	ItemPicUrl      string `json:"item_pic_url" xorm:"not null default '' VARCHAR(255)"`
	Subtitle        string `json:"subtitle" xorm:"not null default '' VARCHAR(255)"`
	TopicCategoryId int    `json:"topic_category_id" xorm:"not null default 0 INT(11)"`
	PriceInfo       string `json:"price_info" xorm:"not null default 0.00 DECIMAL(10,2)"`
	ReadCount       string `json:"read_count" xorm:"not null default '0' VARCHAR(255)"`
	ScenePicUrl     string `json:"scene_pic_url" xorm:"not null default '' VARCHAR(255)"`
	TopicTemplateId int    `json:"topic_template_id" xorm:"not null default 0 INT(11)"`
	TopicTagId      int    `json:"topic_tag_id" xorm:"not null default 0 INT(11)"`
	SortOrder       int    `json:"sort_order" xorm:"not null default 100 INT(11)"`
	IsShow          int    `json:"is_show" xorm:"not null default 1 TINYINT(1)"`
}
