package models

type GoodsIssue struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	GoodsId  string `json:"goods_id" xorm:"TEXT"`
	Question string `json:"question" xorm:"VARCHAR(255)"`
	Answer   string `json:"answer" xorm:"VARCHAR(45)"`
}
