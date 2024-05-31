package model

import (
	"Project/common"
	"errors"
)

var (
	ErrTitleIsEmpty = errors.New("title cannot be empty")
)

type ToDoItem struct {
	common.SqlModel
	Title       string      `gorm:"column:title" json:"title"`
	Description string      `gorm:"column:description" json:"description"`
	Status      *ItemStatus `gorm:"column:status" json:"status"`
}

func (ToDoItem) TableName() string { return "todo_items" }

type CreateToDoItemsPresenter struct {
	Id          int         `gorm:"column:id" json:"_"`
	Title       string      `gorm:"column:title" json:"title"`
	Description string      `gorm:"column:description" json:"description"`
	Status      *ItemStatus `gorm:"column:status" json:"status"`
}

func (CreateToDoItemsPresenter) TableName() string { return ToDoItem{}.TableName() }

type UpdateToDoItemsPresenter struct {
	Title       *string     `gorm:"column:title" json:"title"`
	Description *string     `gorm:"column:description" json:"description"`
	Status      *ItemStatus `gorm:"column:status" json:"status"`
}

func (UpdateToDoItemsPresenter) TableName() string { return ToDoItem{}.TableName() }
