package models

type Address struct {
	Id         int    `json:"id" xorm:"not null pk autoincr MEDIUMINT(8)"`
	Name       string `json:"name" xorm:"not null default '' VARCHAR(50)"`
	UserId     int    `json:"user_id" xorm:"not null default 0 index MEDIUMINT(8)"`
	CountryId  int    `json:"country_id" xorm:"not null default 0 SMALLINT(5)"`
	ProvinceId int    `json:"province_id" xorm:"not null default 0 SMALLINT(5)"`
	CityId     int    `json:"city_id" xorm:"not null default 0 SMALLINT(5)"`
	DistrictId int    `json:"district_id" xorm:"not null default 0 SMALLINT(5)"`
	Address    string `json:"address" xorm:"not null default '' VARCHAR(120)"`
	Mobile     string `json:"mobile" xorm:"not null default '' VARCHAR(60)"`
	IsDefault  int    `json:"is_default" xorm:"not null default 0 TINYINT(1)"`

	CountryName  string `json:"country_name" xorm:"-"`
	ProvinceName string `json:"province_name" xorm:"-"`
	CityName     string `json:"city_name" xorm:"-"`
	FullRegion   string `json:"full_region" xorm:"-"`
}
