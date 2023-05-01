package blog

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/monkeydnoya/hiraishin-blog/pkg/domain"

	configuration "github.com/monkeydnoya/hiraishin-blog/pkg/config"
)

// TODO: Move to domain models
type DBResponse struct {
	ID        string    `json:"id,omitempty"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Role      []string  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func (s Service) GetBlogs() ([]domain.BlogResponse, error) {
	blogList, err := s.DAO.GetBlogs()
	if err != nil {
		configuration.Logger.Errorw("blogs: blogs not found: ", err)
		return nil, err
	}

	return blogList, nil
}

func (s Service) GetBlogById(id string) (domain.BlogResponse, error) {
	blog, err := s.DAO.GetBlogById(id)
	if err != nil {
		configuration.Logger.Errorw("blogs: blog not found: ", err)
		return domain.BlogResponse{}, err
	}

	return blog, nil
}

func (s Service) CreateBlog(blog domain.Blog, user domain.User) (domain.BlogResponse, error) {
	result, err := s.DAO.CreateBlog(blog, user)
	if err != nil {
		configuration.Logger.Errorw("blogs: blog not created: ", err)
		return domain.BlogResponse{}, err
	}

	configuration.Logger.Infow("blogs: new blog was created")
	return result, nil
}

func (s Service) UpdateBlog(blog domain.BlogResponse, user domain.User) (domain.BlogResponse, error) {
	updatedBlog, err := s.DAO.UpdateBlog(blog, user)
	if err != nil {
		configuration.Logger.Errorw("blogs: blog not updated: ", err)
		return domain.BlogResponse{}, err
	}
	return updatedBlog, nil
}

func (s Service) DeleteBlog(id string) error {
	err := s.DAO.DeleteBlog(id)
	if err != nil {
		configuration.Logger.Errorw("blogs: blog not deleted: ", err)
		return err
	}
	return nil
}

func (s Service) DeserializeUserRemote() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var accessToken string = ctx.Cookies("access_token")

		a := fiber.AcquireAgent()
		a.Cookie("access_token", accessToken)
		defer fiber.ReleaseAgent(a)
		req := a.Request()
		req.SetRequestURI(configuration.Config("AUTH_HOST") + "/api/auth/validate")
		req.Header.SetMethod(fiber.MethodGet)
		req.Header.Set("Content-Type", "application/json")

		a.Parse()
		_, b, _ := a.Bytes()

		var user domain.UserValidated
		err := json.Unmarshal(b, &user)
		if err != nil {
			configuration.Logger.Error("user unauthorized: ", err)
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		ctx.Locals("currentUser", user)
		return ctx.Next()
	}
}
