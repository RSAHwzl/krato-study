package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/app/user/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo userRepo) CreateUser(ctx context.Context, user *biz.User) *biz.User {
	repo.data.MysqlDb.Create(&user)
	return user
}

func (repo userRepo) GetUserList(ctx context.Context) ([]*biz.User, error) {
	var users []*biz.User
	err := repo.data.MysqlDb.Model(&biz.User{}).Find(&users).Error
	return users, err
}
