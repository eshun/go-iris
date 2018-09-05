package models

type OrderGoods struct {
	Id                        int    `json:"id" xorm:"not null pk autoincr MEDIUMINT(8)"`
	OrderId                   int    `json:"order_id" xorm:"not null default 0 index MEDIUMINT(8)"`
	GoodsId                   int    `json:"goods_id" xorm:"not null default 0 index MEDIUMINT(8)"`
	GoodsName                 string `json:"goods_name" xorm:"not null default '' VARCHAR(120)"`
	GoodsSn                   string `json:"goods_sn" xorm:"not null default '' VARCHAR(60)"`
	ProductId                 int    `json:"product_id" xorm:"not null default 0 MEDIUMINT(8)"`
	Number                    int    `json:"number" xorm:"not null default 1 SMALLINT(5)"`
	MarketPrice               string `json:"market_price" xorm:"not null default 0.00 DECIMAL(10,2)"`
	RetailPrice               string `json:"retail_price" xorm:"not null default 0.00 DECIMAL(10,2)"`
	GoodsSpecifitionNameValue string `json:"goods_specifition_name_value" xorm:"not null TEXT"`
	IsReal                    int    `json:"is_real" xorm:"not null default 0 TINYINT(1)"`
	GoodsSpecifitionIds       string `json:"goods_specifition_ids" xorm:"not null default '' VARCHAR(255)"`
	ListPicUrl                string `json:"list_pic_url" xorm:"not null default '' VARCHAR(255)"`
}
