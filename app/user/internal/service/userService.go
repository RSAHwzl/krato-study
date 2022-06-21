package service

import (
	"context"
	"errors"
	user "helloworld/api/user"
	"helloworld/app/user/internal/biz"
)

type UserService struct {
	user.UnimplementedUserServiceServer
	uc *biz.UserUsecase
}

// NewUserService new a user service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserRes, error) {
	res := s.uc.CreateUser(ctx, req.Name, req.Password)
	if res.ID == 0 {
		return &user.CreateUserRes{Id: 0}, errors.New("创建失败")
	}

	return &user.CreateUserRes{Id: int32(res.ID)}, nil
}

func (s *UserService) GetUserList(ctx context.Context, req *user.GetUserListReq) (*user.GetUserListRes, error) {
	res := make([]*user.GetUserListRes_User, 0)

	users, err := s.uc.GetUserList(ctx)
	if err != nil {
		return &user.GetUserListRes{Users: res}, err
	}
	for _, v := range users {
		res = append(res, &user.GetUserListRes_User{
			Id:       int32(v.ID),
			Name:     v.Name,
			Password: v.Password,
		})
	}

	return &user.GetUserListRes{Users: res}, nil
}

func (s *UserService) SetUserCache(ctx context.Context, req *user.SetUserCacheReq) (*user.SetUserCacheRes, error) {
	var u = &biz.UserCache{
		Name:    req.Name,
		Age:     req.Age,
		Sex:     req.Sex,
		Address: req.Address,
	}
	res := s.uc.SetUserCache(ctx, u)
	if res == 0 {
		return &user.SetUserCacheRes{Id: 0}, errors.New("缓存失败")
	}

	return &user.SetUserCacheRes{Id: int32(res)}, nil
}

func (s *UserService) GetUserCache(ctx context.Context, req *user.GetUserCacheReq) (*user.GetUserCacheRes, error) {
	res := make(map[string]*user.GetUserCacheRes_User, 0)

	users, err := s.uc.GetUserCache(ctx, req.Name)
	if err != nil {
		return &user.GetUserCacheRes{}, errors.New("缓存获取失败")
	}
	for _, v := range users {
		res[v.Name] = &user.GetUserCacheRes_User{
			Name:    v.Name,
			Age:     v.Age,
			Sex:     v.Sex,
			Address: v.Address,
		}
	}

	return &user.GetUserCacheRes{Users: res}, nil
}
