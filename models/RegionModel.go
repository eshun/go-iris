package models

type Region struct {
	Id       int    `json:"id" xorm:"not null pk autoincr SMALLINT(5)"`
	ParentId int    `json:"parent_id" xorm:"not null default 0 index SMALLINT(5)"`
	Name     string `json:"name" xorm:"not null default '' VARCHAR(120)"`
	Type     int    `json:"type" xorm:"not null default 2 index TINYINT(1)"`
	AgencyId int    `json:"agency_id" xorm:"not null default 0 index SMALLINT(5)"`
}
