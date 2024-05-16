package baserouter

import (
	"insider_case_study/cmd/basecontroller"
	"insider_case_study/db/connection"
	messageRouter "insider_case_study/pkg/message/router"

	"github.com/gofiber/fiber/v2"
)

func InitializeRouters(app fiber.Router, client connection.Client) {

	// API v1 routes
	api := app.Group("/v1")

	// Health check routes
	basecontroller.HealthCheck(api)
	basecontroller.SwaggerController(api)

	messageRouter.InitializeRouter(api, client)

}
