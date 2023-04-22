package mongodriver

import (
	"github.com/monkeydnoya/hiraishin-blog/internal/data/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Client   *mongo.Client
	Database *mongo.Database
}

type BlogDAO struct {
	DB *mongo.Database
}

func (c Config) Init() (repository.BlogRepository, error) {
	blog := BlogDAO{
		DB: c.Database,
	}

	return blog, nil
}
