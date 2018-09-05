package models

type Category struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name         string `json:"name" xorm:"not null default '' VARCHAR(90)"`
	Keywords     string `json:"keywords" xorm:"not null default '' VARCHAR(255)"`
	FrontDesc    string `json:"front_desc" xorm:"not null default '' VARCHAR(255)"`
	ParentId     int    `json:"parent_id" xorm:"not null default 0 index INT(11)"`
	SortOrder    int    `json:"sort_order" xorm:"not null default 50 TINYINT(1)"`
	ShowIndex    int    `json:"show_index" xorm:"not null default 0 TINYINT(1)"`
	IsShow       int    `json:"is_show" xorm:"not null default 1 TINYINT(1)"`
	BannerUrl    string `json:"banner_url" xorm:"not null default '' VARCHAR(255)"`
	IconUrl      string `json:"icon_url" xorm:"not null VARCHAR(255)"`
	ImgUrl       string `json:"img_url" xorm:"not null VARCHAR(255)"`
	WapBannerUrl string `json:"wap_banner_url" xorm:"not null VARCHAR(255)"`
	Level        string `json:"level" xorm:"not null VARCHAR(255)"`
	Type         int    `json:"type" xorm:"not null default 0 INT(11)"`
	FrontName    string `json:"front_name" xorm:"not null VARCHAR(255)"`

	GoodsList []*Goods `json:"-" xorm:"-"`
	SubCategoryList []*Category `json:"subCategoryList" xorm:"-"`
}
