package handler

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/monkeydnoya/hiraishin-blog/pkg/domain"

	configuration "github.com/monkeydnoya/hiraishin-blog/pkg/config"
)

func (s Server) GetBlogs(c *fiber.Ctx) error {
	blogList, err := s.Service.GetBlogs()
	if err != nil {
		return c.Status(404).JSON(err)
	}

	return c.Status(200).JSON(blogList)
}

func (s Server) GetBlogById(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		configuration.Logger.Errorw("incorrect id format:")
		return c.SendStatus(400)
	}

	blog, err := s.Service.GetBlogById(id)
	if err != nil {
		return c.Status(404).JSON(err)
	}

	return c.Status(200).JSON(blog)
}

func (s Server) CreateBlog(c *fiber.Ctx) error {
	var blog domain.Blog
	var user domain.User

	err := c.BodyParser(&blog)
	if err != nil {
		configuration.Logger.Errorw("can't parse body:", err)
		return c.SendStatus(400)
	}

	data := c.Locals("currentUser")

	dataString := fmt.Sprintf("%s", data)
	dataFields := strings.Split(dataString, " ")

	user.Username = dataFields[3]
	user.Email = dataFields[4]

	result, err := s.Service.CreateBlog(blog, user)
	if err != nil {
		return c.SendStatus(400)
	}

	return c.Status(200).JSON(result)
}

func (s Server) UpdateBlog(c *fiber.Ctx) error {
	var blogToUpdate domain.BlogResponse
	var user domain.User

	err := c.BodyParser(&blogToUpdate)
	if err != nil {
		configuration.Logger.Errorw("can't parse body:", err)
		return c.SendStatus(400)
	}

	user.Username = c.Cookies("username")
	user.Email = c.Cookies("email")

	result, err := s.Service.UpdateBlog(blogToUpdate, user)
	if err != nil {
		return c.SendStatus(400)
	}
	return c.Status(200).JSON(result)
}

func (s Server) DeleteBlog(c *fiber.Ctx) error {
	id, err := url.PathUnescape(c.Path("id"))
	if err != nil {
		configuration.Logger.Errorw("can't parse configuration from path:", err)
		return c.SendStatus(400)
	}

	if err := s.Service.DeleteBlog(id); err != nil {
		return c.SendStatus(400)
	}
	return c.SendStatus(200)
}

func (s Server) DeserializeUserRemote() fiber.Handler {
	return s.Service.DeserializeUserRemote()
}
