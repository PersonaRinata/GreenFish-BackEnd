package dao

import (
	"GreenFish/server/common/consts"
	"GreenFish/server/service/user/model"
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func (u User) CreateUser(ctx context.Context, user *model.User) error {
	var temp model.User
	err := u.db.Where("username = ?", user.Username).First(&temp).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		klog.Error("mysql select failed,", err)

		return err
	}
	if temp.Username != "" {
		err = errors.New(consts.MysqlAlreadyExists)
		return err
	}
	err = u.db.Create(&user).Error
	if err != nil {
		klog.Error("mysql insert failed", err)

		return err
	}

	return nil
}

func (u User) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func (u User) SearchUserByUsername(ctx context.Context, content string) ([]*model.User, error) {
	var users []*model.User
	content = "%" + content + "%"
	if err := u.db.Where("username LIKE ?", content).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u User) ChangeAvatarByUserID(ctx context.Context, avatar string, id int64) error {
	err := u.db.Model(&model.User{}).WithContext(ctx).Where("id = ?", id).Update("avatar", avatar).Error
	if err != nil {
		return err
	}
	return err
}

func (u User) ChangeNicknameByUserID(ctx context.Context, nickname string, id int64) error {
	err := u.db.Model(&model.User{}).WithContext(ctx).Where("id = ?", id).Update("nickname", nickname).Error
	if err != nil {
		return err
	}
	return err
}

func (u User) AddDoctor(ctx context.Context, id int64, department string) error {
	if err := u.db.Where("id = ?", id).Update("department = ?", department).Error; err != nil {

		return err
	}

	return nil
}

func NewUser(db *gorm.DB) *User {
	m := db.Migrator()
	if !m.HasTable(&model.User{}) {
		err := m.CreateTable(&model.User{})
		if err != nil {
			klog.Error("create mysql table failed,", err)
		}
	}
	return &User{db: db}
}
