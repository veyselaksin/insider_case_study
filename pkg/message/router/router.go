package router

import (
	"github.com/gofiber/fiber/v2"
	"insider_case_study/db/connection"
	"insider_case_study/db/repository/userrepo"
	"insider_case_study/pkg/message/controller"
	"insider_case_study/pkg/message/service"
)

func InitializeRouter(router fiber.Router, client connection.Client) {
	userRepository := userrepo.New(client.PostgresConnection)
	userService := service.New(userRepository)
	messageController := controller.New(userService)

	// Message routes
	endpoint := router.Group("/message")
	endpoint.Post("/", messageController.SendMessage)
}
