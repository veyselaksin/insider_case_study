package controller

import (
	"github.com/gofiber/fiber/v2"
	"insider_case_study/cmd/baseresponse"
	"insider_case_study/dto/messagedto"
	"insider_case_study/pkg/message/service"
)

type Controller interface {
	SendMessage(ctx *fiber.Ctx) error
	MessageRecipients(ctx *fiber.Ctx) error
}

type controller struct {
	service service.Service
}

func New(service service.Service) Controller {
	return &controller{
		service: service,
	}
}

// SendMessage Message godoc
// @Summary Send message
// @Description This endpoint is used to send message every 2 minutes when tickerStatus parameter is start
// @Tags Message
// @Accept json
// @Produce json
// @Param sendMessageRequest body messagedto.SendMessageRequest true "Send message requests"
// @Success 200 {object} map[string]interface{}
// @Router /message/ [post]
func (c *controller) SendMessage(ctx *fiber.Ctx) error {
	var request messagedto.SendMessageRequest
	if err := ctx.BodyParser(&request); err != nil {
		return baseresponse.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	err := c.service.SendMessage(ctx, request)
	if err != nil {
		return baseresponse.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return baseresponse.SuccessResponse(ctx, nil, fiber.StatusOK)
}

// MessageRecipients godoc
// @Summary Get message recipients
// @Description This endpoint is used to get all message recipients
// @Tags Message
// @Accept json
// @Produce json
// @Success 200 {object} []messagedto.MessageRecipientsResponse
// @Router /message/recipients [get]
func (c *controller) MessageRecipients(ctx *fiber.Ctx) error {
	data, err := c.service.MessageRecipients(ctx)
	if err != nil {
		return baseresponse.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return baseresponse.SuccessResponse(ctx, data, fiber.StatusOK)
}
