package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monkeydnoya/hiraishin-blog/pkg/domain"
)

type BlogService interface {
	GetBlogs() ([]domain.BlogResponse, error)
	GetBlogById(id string) (domain.BlogResponse, error)
	CreateBlog(domain.Blog, domain.User) (domain.BlogResponse, error)
	UpdateBlog(domain.BlogResponse, domain.User) (domain.BlogResponse, error)
	DeleteBlog(id string) error

	DeserializeUserRemote() fiber.Handler
}
