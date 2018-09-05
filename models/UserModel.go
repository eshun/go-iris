package models

type User struct {
	Id            int    `json:"id" xorm:"not null pk autoincr MEDIUMINT(8)"`
	Username      string `json:"username" xorm:"not null default '' unique VARCHAR(60)"`
	Password      string `json:"-" xorm:"not null default '' VARCHAR(32)"`
	Gender        int    `json:"gender" xorm:"not null default 0 TINYINT(1)"`
	Birthday      int    `json:"birthday" xorm:"not null default 0 INT(11)"`
	RegisterTime  int    `json:"register_time" xorm:"not null default 0 INT(11)"`
	LastLoginTime int    `json:"last_login_time" xorm:"not null default 0 INT(11)"`
	LastLoginIp   string `json:"last_login_ip" xorm:"not null default '' VARCHAR(15)"`
	UserLevelId   int    `json:"user_level_id" xorm:"not null default 0 TINYINT(3)"`
	Nickname      string `json:"nickname" xorm:"not null VARCHAR(60)"`
	Mobile        string `json:"mobile" xorm:"not null VARCHAR(20)"`
	RegisterIp    string `json:"register_ip" xorm:"not null default '' VARCHAR(45)"`
	Avatar        string `json:"avatar" xorm:"not null default '' VARCHAR(255)"`
	WeixinOpenid  string `json:"weixin_openid" xorm:"not null default '' VARCHAR(50)"`
}
//func (u *User) TableName() string {
//	return "user"
//}
//
//func (u *User) Valid(v *validation.Validation) {
//
//}