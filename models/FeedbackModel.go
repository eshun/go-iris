package models

type Feedback struct {
	MsgId      int    `json:"msg_id" xorm:"not null pk autoincr MEDIUMINT(8)"`
	ParentId   int    `json:"parent_id" xorm:"not null default 0 MEDIUMINT(8)"`
	UserId     int    `json:"user_id" xorm:"not null default 0 index MEDIUMINT(8)"`
	UserName   string `json:"user_name" xorm:"not null default '' VARCHAR(60)"`
	UserEmail  string `json:"user_email" xorm:"not null default '' VARCHAR(60)"`
	MsgTitle   string `json:"msg_title" xorm:"not null default '' VARCHAR(200)"`
	MsgType    int    `json:"msg_type" xorm:"not null default 0 TINYINT(1)"`
	MsgStatus  int    `json:"msg_status" xorm:"not null default 0 TINYINT(1)"`
	MsgContent string `json:"msg_content" xorm:"not null TEXT"`
	MsgTime    int    `json:"msg_time" xorm:"not null default 0 INT(10)"`
	MessageImg string `json:"message_img" xorm:"not null default '0' VARCHAR(255)"`
	OrderId    int    `json:"order_id" xorm:"not null default 0 INT(11)"`
	MsgArea    int    `json:"msg_area" xorm:"not null default 0 TINYINT(1)"`
}
