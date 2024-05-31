package biz

import (
	"Project/modules/item/model"
	"context"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.CreateToDoItemsPresenter) error
}

type CreateItemBiz struct {
	storage CreateItemStorage
}

func NewCreateItemBiz(storage CreateItemStorage) *CreateItemBiz {
	return &CreateItemBiz{storage: storage}
}

func (biz *CreateItemBiz) CreateNewItem(ctx context.Context, data *model.CreateToDoItemsPresenter) error {
	title := strings.TrimSpace(data.Title)
	if title == "" {
		return model.ErrTitleIsEmpty
	}
	if err := biz.storage.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
