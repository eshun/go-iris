package models

type Attribute struct {
	Id                  int    `xorm:"not null pk autoincr INT(11)"`
	AttributeCategoryId int    `xorm:"not null default 0 index INT(11)"`
	Name                string `xorm:"not null default '' VARCHAR(60)"`
	InputType           int    `xorm:"not null default 1 TINYINT(1)"`
	Values              string `xorm:"not null TEXT"`
	SortOrder           int    `xorm:"not null default 0 TINYINT(3)"`
}
