package mongo

import (
	"context"
	"github.com/daycolor/chat-app/config"
	"github.com/daycolor/chat-app/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type User struct {
	*models.User
	Token string
}

func ConnectDB() (*mongo.Database, error) {
	mongoConfig := config.MongoConfig{}
	mongoConfig.Read()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConfig.ConnectionUrl))

	if err != nil {
		return nil, err
	}

	return client.Database("chat"), nil
}
