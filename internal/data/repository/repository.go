package repository

import "github.com/monkeydnoya/hiraishin-blog/pkg/domain"

type BlogRepository interface {
	GetBlogs() ([]domain.BlogResponse, error)
	GetBlogById(id string) (domain.BlogResponse, error)
	CreateBlog(domain.Blog, domain.User) (domain.BlogResponse, error)
	UpdateBlog(domain.BlogResponse, domain.User) (domain.BlogResponse, error) // Refactor taked argument, this gonna be new blog struct, not named Response
	DeleteBlog(id string) error
}
