package test

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	user "helloworld/api/user"
	"testing"
)

func TestGrpc(t *testing.T) {
	ctx := context.Background()
	con, err := grpc.DialInsecure(ctx, grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer con.Close()

	service := user.NewUserServiceClient(con)
	res, err := service.GetUserList(ctx, &user.GetUserListReq{Name: "789"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
