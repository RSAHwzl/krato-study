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
