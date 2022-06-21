package cacheMongo

import (
	"context"
	"helloworld/app/user/internal/conf"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoRepo(c *conf.Data) (context.Context, *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.Mongo.Source))
	if err != nil {
		log.Println("failed to connect dbMysql:", err)
		panic(err)
	}

	return ctx, client
}
