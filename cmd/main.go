package main

import (
	"log"

	"github.com/monkeydnoya/hiraishin-blog/internal/controller"
	"github.com/monkeydnoya/hiraishin-blog/internal/controller/fiber/handler"
	"github.com/monkeydnoya/hiraishin-blog/internal/data/config/mongodb"
	"github.com/monkeydnoya/hiraishin-blog/internal/data/repository/mongodriver"
	"github.com/monkeydnoya/hiraishin-blog/internal/service/blog"
	"github.com/monkeydnoya/hiraishin-blog/pkg/config"
)

func main() {
	config := &mongodb.Config{
		Host:     config.Config("MONGODB_URI"),
		Port:     config.Config("MONGODB_PORT"),
		DBName:   config.Config("MONGODB_DBNAME"),
		Username: config.Config("MONGODB_USERNAME"),
		Password: config.Config("MONGODB_PASSWORD"),
	}
	connection, err := config.Connect()
	// Rethink: Error handling is correct?
	// If os.Exit executes in config.Connect()
	if err != nil {
		log.Fatal(err)
	}

	blogDAO, err := mongodriver.Config{Client: connection.Client, Database: connection.DbConnection}.Init()
	// TODO: Add correctly error handling
	if err != nil {
		log.Fatal(err)
	}

	service, err := blog.NewService(blog.Service{DAO: blogDAO})
	// TODO: Add correctly error handling
	if err != nil {
		log.Fatal(err)
	}

	server := handler.NewServer(service)
	controller.StartServer(server)
}
