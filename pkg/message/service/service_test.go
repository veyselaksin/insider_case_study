package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"go.uber.org/mock/gomock"
	"insider_case_study/db/models"
	"insider_case_study/dto/messagedto"
	"insider_case_study/i18n"
	"insider_case_study/mocks/repository/userrepo"
	"testing"
)

var fiberCtx *fiber.Ctx

var s Service

var mockUserRepository *userrepo.MockUserRepository

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	app := fiber.New()
	fiberCtx = app.AcquireCtx(&fasthttp.RequestCtx{})

	// Assign language to fiber context header
	fiberCtx.Request().Header.Set("Accept-Language", "en")

	i18n.InitBundle("./../../../i18n/languages")

	mockUserRepository = userrepo.NewMockUserRepository(ct)

	s = New(mockUserRepository)
	return func() {
		s = nil
		defer ct.Finish()
	}
}

func TestService_SendMessage_WhenUserRepoFindAllError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockUserRepository.EXPECT().FindAll(gomock.Any()).Return(nil, errors.New("Record not found."))

	err := s.SendMessage(fiberCtx, messagedto.SendMessageRequest{})

	assert.Error(t, err)
	assert.Equal(t, "An unexpected error occurred.", err.Error())
	assert.Nil(t, fiberCtx.Response().Body())
}

func TestService_SendMessage_WhenTickerStatusIsStart(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockUserRepository.EXPECT().FindAll(gomock.Any()).Return([]models.User{
		{
			ID:          "123e4567-e89b-12d3-a456-426614174000",
			PhoneNumber: "1234567890",
			Status:      "pending",
		},
	}, nil)

	err := s.SendMessage(fiberCtx, messagedto.SendMessageRequest{
		TickerStatus: "start",
	})

	assert.NoError(t, err)
}

func TestService_SendMessage_WhenTickerStatusIsStop(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockUserRepository.EXPECT().FindAll(gomock.Any()).Return([]models.User{
		{
			ID:          "123e4567-e89b-12d3-a456-426614174000",
			PhoneNumber: "1234567890",
			Status:      "pending",
		},
	}, nil)

	err := s.SendMessage(fiberCtx, messagedto.SendMessageRequest{
		TickerStatus: "stop",
	})

	assert.NoError(t, err)
}

func TestService_SendMessage_WhenTickerStatusIsInvalid(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockUserRepository.EXPECT().FindAll(gomock.Any()).Return([]models.User{
		{
			ID:          "123e4567-e89b-12d3-a456-426614174000",
			PhoneNumber: "1234567890",
			Status:      "pending",
		},
	}, nil)

	err := s.SendMessage(fiberCtx, messagedto.SendMessageRequest{
		TickerStatus: "invalid",
	})

	assert.Nil(t, err)
}

func TestService_SendMessage_WhenRequestTickerStatusIsInvalid(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockUserRepository.EXPECT().FindAll(gomock.Any()).Return([]models.User{
		{
			ID:          "123e4567-e89b-12d3-a456-426614174000",
			PhoneNumber: "1234567890",
			Status:      "pending",
		},
	}, nil)

	err := s.SendMessage(fiberCtx, messagedto.SendMessageRequest{})

	assert.Nil(t, err)
}

func TestService_MessageRecipients_WhenUserRepoFindAllError(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockUserRepository.EXPECT().FindAll(gomock.Any()).Return(nil, errors.New("Record not found."))

	_, err := s.MessageRecipients(fiberCtx)

	assert.Error(t, err)
	assert.Equal(t, "Record not found.", err.Error())
}

func TestService_MessageRecipients_WhenUserRepoFindAllSuccess(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockUserRepository.EXPECT().FindAll(gomock.Any()).Return([]models.User{
		{
			ID:          "123e4567-e89b-12d3-a456-426614174000",
			PhoneNumber: "1234567890",
			Status:      "sent",
		},
		{
			ID:          "123e4567-e89b-12d3-a456-426614174001",
			PhoneNumber: "1234567891",
			Status:      "sent",
		},
	}, nil)

	response, err := s.MessageRecipients(fiberCtx)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", response[0].ID)
	assert.Equal(t, "1234567890", response[0].PhoneNumber)
	assert.Equal(t, "sent", response[0].Status)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174001", response[1].ID)
	assert.Equal(t, "1234567891", response[1].PhoneNumber)
	assert.Equal(t, "sent", response[1].Status)
}

func TestService_MessageRecipients_WhenUserRepoFindAllSuccessButNoData(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockUserRepository.EXPECT().FindAll(gomock.Any()).Return(nil, nil)

	response, err := s.MessageRecipients(fiberCtx)

	assert.NoError(t, err)
	assert.Nil(t, response)
}
