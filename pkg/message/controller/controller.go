package controller

import (
	"github.com/gofiber/fiber/v2"
	"insider_case_study/cmd/baseresponse"
	"insider_case_study/dto/messagedto"
	"insider_case_study/pkg/message/service"
)

type Controller interface {
	SendMessage(ctx *fiber.Ctx) error
}

type controller struct {
	service service.Service
}

func New(service service.Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) SendMessage(ctx *fiber.Ctx) error {
	var request messagedto.SendMessageRequest
	if err := ctx.BodyParser(&request); err != nil {
		return baseresponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	err := c.service.SendMessage(request)
	if err != nil {
		return baseresponse.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return baseresponse.SuccessResponse(ctx, nil, fiber.StatusOK)
}
