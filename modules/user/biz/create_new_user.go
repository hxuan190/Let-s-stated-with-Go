package biz

import (
	"Project/common"
	"Project/modules/user/model"
	"context"
	"strings"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *model.Register) error
	FindUser(ctx context.Context, email string) bool
}
type CreateUserBiz struct {
	storage CreateUserStorage
}

func NewCreateUserBiz(storage CreateUserStorage) *CreateUserBiz {
	return &CreateUserBiz{}
}

func (biz *CreateUserBiz) CreateNewUser(ctx context.Context, data *model.Register) error {
	if firstname := strings.TrimSpace(data.FirstName); firstname == "" {
		return model.ErrFirstNameIsEmpty
	} else if lastname := strings.TrimSpace(data.LastName); lastname == "" {
		return model.ErrLastNameIsEmpty
	} else if email := strings.TrimSpace(data.Email); email == "" {
		return model.ErrEmailIsEmpty
	} else if pw := strings.TrimSpace(data.Password); pw == "" {
		return model.ErrPassWordIsEmpty
	}

	if isUser := biz.storage.FindUser(ctx, data.Email); isUser == true {
		return model.ErrUserIsExist
	}
	salt, err := common.GenerateRandomSalt(16)
	if err != nil {
		return err
	}
	hashedPassword, err := common.HashPasswordWithSalt(data.Password, salt)
	if err != nil {
		return err
	}
	data.Password = hashedPassword

	if err := biz.storage.CreateUser(ctx, data); err != nil {
		return err
	}
	return nil
}
