package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
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

func (repo userRepo) SetUserCache(ctx context.Context, u *biz.UserCache) int {
	//ct, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	collection := repo.data.Mongo.Database("demo").Collection("student")
	_, err := collection.InsertOne(ctx, bson.D{
		{"name", u.Name},
		{"age", u.Age},
		{"sex", u.Sex},
		{"address", u.Address},
	})
	if err != nil {
		return 0
	}

	return 1
}

func (repo userRepo) GetUserCache(ctx context.Context, name string) ([]*biz.UserCache, error) {
	result := make([]*biz.UserCache, 0)
	//ct, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	collection := repo.data.Mongo.Database("demo").Collection("student")
	cur, err := collection.Find(ctx, bson.D{{"name", name}})
	if err != nil {
		repo.log.Error(err)
		return result, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var tmp *biz.UserCache
		err = cur.Decode(&tmp)
		if err != nil {
			repo.log.Error(err)
			return result, err
		}
		result = append(result, tmp)
	}
	if err = cur.Err(); err != nil {
		repo.log.Error(err)
		return result, err
	}

	return result, nil
}
