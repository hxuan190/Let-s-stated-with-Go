package common

import (
	"github.com/google/uuid"
	"time"
)

type SQLModel struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
type SqlUserModel struct {
	Id        uuid.UUID  `gorm:"column:id" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

func NewSQLModel() SQLModel {
	id, _ := uuid.NewV7()
	now := time.Now().UTC()

	return SQLModel{
		ID:        id,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
