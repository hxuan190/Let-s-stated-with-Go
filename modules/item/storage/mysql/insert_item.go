package mysql

import (
	"Project/modules/item/model"
	"context"
)

func (s *sqlStorage) CreateItem(ctx context.Context, data *model.CreateToDoItemsPresenter) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
