package models

type Brand struct {
	Id            int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name          string `json:"name" xorm:"not null default '' VARCHAR(255)"`
	ListPicUrl    string `json:"list_pic_url" xorm:"not null default '' VARCHAR(255)"`
	SimpleDesc    string `json:"simple_desc" xorm:"not null default '' VARCHAR(255)"`
	PicUrl        string `json:"pic_url" xorm:"not null default '' VARCHAR(255)"`
	SortOrder     int    `json:"sort_order" xorm:"not null default 50 TINYINT(3)"`
	IsShow        int    `json:"is_show" xorm:"not null default 1 index TINYINT(1)"`
	FloorPrice    string `json:"floor_price" xorm:"not null default 0.00 DECIMAL(10,2)"`
	AppListPicUrl string `json:"app_list_pic_url" xorm:"not null default '' VARCHAR(255)"`
	IsNew         int    `json:"is_new" xorm:"not null default 0 TINYINT(1)"`
	NewPicUrl     string `json:"new_pic_url" xorm:"not null default '' VARCHAR(255)"`
	NewSortOrder  int    `json:"new_sort_order" xorm:"not null default 10 TINYINT(2)"`
}
