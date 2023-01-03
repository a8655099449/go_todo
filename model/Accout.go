package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);column(username);not null;unique;comment: " json:"nickname,omitempty"`
	Password string `gorm:"type:varchar(200);not null;comment:密码"`
	Nickname string `gorm:"type:varchar(50);not null;comment:用户昵称"`
}

func (Account) TableName() string {
	return "go_user"
}
