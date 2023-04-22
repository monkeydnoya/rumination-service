package controller

import (
	"os"

	"github.com/monkeydnoya/hiraishin-blog/internal/controller/fiber/handler"
	configuration "github.com/monkeydnoya/hiraishin-blog/pkg/config"
)

func StartServer(server handler.Server) {
	configuration.Logger.Infow("Starting Blog Service")

	os.Getenv(".env")
	port := ":" + configuration.Config("SERVICE_PORT")

	server.App.Listen(port)
}
