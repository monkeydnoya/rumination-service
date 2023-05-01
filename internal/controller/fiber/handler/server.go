package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/monkeydnoya/hiraishin-blog/internal/service"
)

type Server struct {
	Service service.BlogService
	App     *fiber.App
}

func NewServer(service service.BlogService) Server {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, X-Api-Key, X-Requested-With, Content-Type, Accept, Authorization, authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	server := Server{
		Service: service,
		App:     app,
	}

	server.SetupRoutes()
	return server
}
