package biz

import (
	"context"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserCache struct {
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	Sex     int32  `json:"sex"`
	Address string `json:"address"`
}

type UserRepo interface {
	CreateUser(context.Context, *User) *User
	GetUserList(context.Context) ([]*User, error)
	SetUserCache(context.Context, *UserCache) int
	GetUserCache(context.Context, string) ([]*UserCache, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, name, password string) *User {
	user := &User{
		Name:     name,
		Password: password,
	}
	uc.repo.CreateUser(ctx, user)
	return user
}

func (uc *UserUsecase) GetUserList(ctx context.Context) ([]*User, error) {
	return uc.repo.GetUserList(ctx)
}

func (uc *UserUsecase) SetUserCache(ctx context.Context, u *UserCache) int {
	return uc.repo.SetUserCache(ctx, u)
}

func (uc *UserUsecase) GetUserCache(ctx context.Context, name string) ([]*UserCache, error) {
	return uc.repo.GetUserCache(ctx, name)
}
