package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"helloworld/app/user/internal/conf"
	"helloworld/app/user/internal/data/cacheMongo"
	"helloworld/app/user/internal/data/dbMysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserRepo)

// Data .wrapped database client
type Data struct {
	MysqlDb *gorm.DB
	Mongo   *mongo.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	mysqlDb := dbMysql.NewMysqlRepo(c)
	ctx, mongoDb := cacheMongo.NewMongoRepo(c)
	d := &Data{MysqlDb: mysqlDb, Mongo: mongoDb}

	return d, func() {
		//closing the data resources
		if err := mongoDb.Disconnect(ctx); err != nil {
			log.NewHelper(logger).Error("failed to disconnect mongo")
			panic(err)
		}
	}, nil
}
