package model

import (
	"Project/common"
	"errors"
)

var (
	ErrUserIsExist      = errors.New("user is exist")
	ErrFirstNameIsEmpty = errors.New("First Name cannot be empty")
	ErrLastNameIsEmpty  = errors.New("Last Name cannot be empty")
	ErrPassWordIsEmpty  = errors.New("Password cannot be empty")
	ErrEmailIsEmpty     = errors.New("Email cannot be empty")
)

type User struct {
	common.SqlUserModel
	FirstName string   `gorm:"column:first_name" json:"first_name"`
	LastName  string   `gorm:"column:last_name" json:"last_name"`
	Email     string   `gorm:"column:email" json:"email"`
	Password  string   `gorm:"column:pass_word" json:"pass_word"`
	Salt      string   `gorm:"column:salt" json:"salt"`
	Role      UserRole `gorm:"column:role" json:"role"`
}

func (User) TableName() string { return "todo_db" }

type Register struct {
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:pass_word" json:"pass_word"`
}

func (Register) TableName() string { return User{}.TableName() }
