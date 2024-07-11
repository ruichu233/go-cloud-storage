package store

import (
	"context"
	"go-cloud-storage/app/user/internal/model"
	"go-cloud-storage/app/user/internal/store/mysql"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore() *UserStore {
	return &UserStore{mysql.NewDBClient()}
}

func (u *UserStore) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).Model(&user).Where("username = ?", username).First(&user).Error
	return &user, err
}

func (u *UserStore) ExistUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	err := u.db.WithContext(ctx).Model(&model.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

func (u *UserStore) Create(ctx context.Context, user *model.User) error {
	if err := u.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}
