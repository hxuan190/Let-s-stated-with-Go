package mysql

import (
	"Project/modules/user/model"
	"context"
)

func (s *sqlStorage) CreateUser(ctx context.Context, data *model.Register) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
func (s *sqlStorage) FindUser(ctx context.Context, email string) bool {
	if err := s.db.Find(email).Error; err != nil {
		return true
	}
	return false
}
