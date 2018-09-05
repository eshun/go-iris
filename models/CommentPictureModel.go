package models

type CommentPicture struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	CommentId int    `json:"comment_id" xorm:"not null default 0 INT(11)"`
	PicUrl    string `json:"pic_url" xorm:"not null default '' VARCHAR(255)"`
	SortOrder int    `json:"sort_order" xorm:"not null default 5 TINYINT(1)"`
}
