package mongodb

import (
	"context"
	"log"
	"time"

	configuration "github.com/monkeydnoya/hiraishin-blog/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Host     string
	Port     string
	DBName   string
	AppEnv   string
	Username string
	Password string
}

type Repository struct {
	Client       *mongo.Client
	DbConnection *mongo.Database
}

func (config *Config) Connect() (Repository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Host).SetAuth(
		options.Credential{
			Username: config.Username,
			Password: config.Password,
		}))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	configuration.Logger.Info("MongoDB successfully connected")
	return Repository{Client: client, DbConnection: client.Database(config.DBName)}, nil
}
