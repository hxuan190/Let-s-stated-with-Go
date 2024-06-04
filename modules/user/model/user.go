package model

import "Project/common"

type User struct {
	common.SqlModel
	FirstName string   `gorm:"column:first_name" json:"first_name"`
	LastName  string   `gorm:"column:last_name" json:"last_name"`
	Email     string   `gorm:"column:email" json:"email"`
	Password  string   `gorm:"column:pass_word" json:"pass_word"`
	Salt      string   `gorm:"column:salt" json:"salt"`
	Role      UserRole `gorm:"column:role" json:"role"`
}
