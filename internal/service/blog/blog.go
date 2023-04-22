package blog

import (
	"encoding/json"
	"fmt"
	"strings"
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
		var accessToken string
		cookie := ctx.Cookies("access_token")

		authorizationHeader := ctx.Get("Authorization")
		authHeaderString := fmt.Sprint(authorizationHeader)
		fields := strings.Fields(authHeaderString)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if cookie != "" {
			accessToken = cookie
		}
		a := fiber.AcquireAgent()
		defer fiber.ReleaseAgent(a)
		req := a.Request()
		req.SetRequestURI(configuration.Config("AUTH_HOST") + "/api/auth/validate")
		req.Header.Set("Authorization", accessToken)
		req.Header.SetMethod(fiber.MethodGet)
		req.Header.Set("Content-Type", "application/json")

		a.Parse()
		_, b, _ := a.Bytes()

		var valToken domain.ValidateToken
		err := json.Unmarshal(b, &valToken)
		if err != nil {
			configuration.Logger.Error("user unauthorized", err)
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		queryString := fmt.Sprintf("id=%s", valToken.UserID)
		fmt.Println(queryString)

		userAgent := fiber.AcquireAgent()
		userAgent.QueryString(queryString)
		defer fiber.ReleaseAgent(userAgent)
		userReq := userAgent.Request()
		userReq.SetRequestURI(configuration.Config("AUTH_HOST") + "/api/auth/details/" + valToken.UserID)
		userReq.Header.SetMethod(fiber.MethodGet)
		userReq.Header.Set("Content-Type", "application/json")
		userAgent.Parse()

		_, body, _ := userAgent.Bytes()

		var user DBResponse
		err = json.Unmarshal(body, &user)
		if err != nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		ctx.Locals("currentUser", user)
		return ctx.Next()
	}
}

// Rethink: Need to add to Middleware result cookie?
// ctx.Cookie(&fiber.Cookie{
// 	Name:     "email",
// 	Value:    user.Email,
// 	HTTPOnly: true,
// 	SameSite: "lax",
// })
// ctx.Cookie(&fiber.Cookie{
// 	Name:     "username",
// 	Value:    user.UserName,
// 	HTTPOnly: true,
// 	SameSite: "lax",
// })
