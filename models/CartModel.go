package models

type Cart struct {
	Id                        int    `json:"id" xorm:"not null pk autoincr MEDIUMINT(8)"`
	UserId                    int    `json:"user_id" xorm:"not null default 0 MEDIUMINT(8)"`
	SessionId                 string `json:"session_id" xorm:"not null default '' index CHAR(32)"`
	GoodsId                   int    `json:"goods_id" xorm:"not null default 0 MEDIUMINT(8)"`
	GoodsSn                   string `json:"goods_sn" xorm:"not null default '' VARCHAR(60)"`
	ProductId                 int    `json:"product_id" xorm:"not null default 0 MEDIUMINT(8)"`
	GoodsName                 string `json:"goods_name" xorm:"not null default '' VARCHAR(120)"`
	MarketPrice               string `json:"market_price" xorm:"not null default 0.00 DECIMAL(10,2)"`
	RetailPrice               string `json:"retail_price" xorm:"not null default 0.00 DECIMAL(10,2)"`
	Number                    int    `json:"number" xorm:"not null default 0 SMALLINT(5)"`
	GoodsSpecifitionNameValue string `json:"goods_specifition_name_value" xorm:"not null TEXT"`
	GoodsSpecifitionIds       string `json:"goods_specifition_ids" xorm:"not null default '' VARCHAR(60)"`
	Checked                   int    `json:"checked" xorm:"not null default 1 TINYINT(1)"`
	ListPicUrl                string `json:"list_pic_url" xorm:"not null default '' VARCHAR(255)"`
}
