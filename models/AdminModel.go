package models

type Admin struct {
	Id            int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Username      string `json:"username" xorm:"not null default '''' VARCHAR(10)"`
	Password      string `json:"-" xorm:"not null default '''' VARCHAR(255)"`
	PasswordSalt  string `json:"password_salt" xorm:"not null default '''' VARCHAR(255)"`
	LastLoginIp   string `json:"last_login_ip" xorm:"not null default '''' VARCHAR(60)"`
	LastLoginTime int    `json:"last_login_time" xorm:"not null default 0 INT(11)"`
	AddTime       int    `json:"add_time" xorm:"not null default 0 INT(11)"`
	UpdateTime    int    `json:"update_time" xorm:"not null default 0 INT(11)"`
	Avatar        string `json:"avatar" xorm:"not null default '''' VARCHAR(255)"`
	AdminRoleId   int    `json:"admin_role_id" xorm:"not null default 0 INT(11)"`
}
