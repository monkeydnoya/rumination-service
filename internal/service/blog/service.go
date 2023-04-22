package blog

import (
	"github.com/monkeydnoya/hiraishin-blog/internal/data/repository"
	"github.com/monkeydnoya/hiraishin-blog/internal/service"
)

type Service struct {
	DAO repository.BlogRepository
}

func NewService(auth Service) (service.BlogService, error) {
	service := Service{
		DAO: auth.DAO,
	}
	return service, nil
}
