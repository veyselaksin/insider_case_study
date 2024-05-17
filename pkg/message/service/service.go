package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"insider_case_study/db/models"
	"insider_case_study/db/repository/userrepo"
	"insider_case_study/dto/messagedto"
	"insider_case_study/helpers/enums"
	"insider_case_study/helpers/requests"
	"insider_case_study/i18n"
	"insider_case_study/i18n/messages"
	"os"
	"sync"
	"time"
)

type Service interface {
	SendMessage(ctx *fiber.Ctx, request messagedto.SendMessageRequest) error
	MessageRecipients(ctx *fiber.Ctx) ([]messagedto.MessageRecipientsResponse, error)
}

//go:generate mockgen -destination=../../../mocks/services/messageservice/messageservice.go -package=messageservice -source=service.go
type service struct {
	userRepo userrepo.UserRepository
}

func New(userRepo userrepo.UserRepository) Service {
	return &service{
		userRepo: userRepo,
	}
}

var ticker *time.Ticker
var quit chan struct{}
var wg sync.WaitGroup

func (s *service) SendMessage(ctx *fiber.Ctx, request messagedto.SendMessageRequest) error {
	userData, err := s.userRepo.FindAll(map[string]any{"status": "pending"})
	if err != nil {
		log.Error(err)
		return errors.New(i18n.CreateMsg(ctx, messages.UnexpectedError))
	}

	if request.TickerStatus == "start" {
		ticker = time.NewTicker(2 * time.Second)
		quit = make(chan struct{})
		wg.Add(1)

		go func(data []models.User) {
			defer wg.Done()
			for {
				select {
				case <-ticker.C:
					fmt.Println("Ticker is running")
					for _, user := range data {
						resp, err := requests.NewRequest(requests.Request{
							URL:    os.Getenv("WEBHOOK_URL"),
							Method: "POST",
							Body: map[string]string{
								"to":      user.PhoneNumber,
								"content": user.Content,
							},
							Header: map[string]string{"Content-Type": "application/json"},
						})
						if err != nil {
							log.Error(err)
							return
						}
						if resp != nil {
							user.Status = enums.Sent
							err = s.userRepo.Update(&user)
							if err != nil {
								log.Error(err)
								return
							}

							err = s.userRepo.SetMessageResponse(context.Background(), user.ID, string(resp))
							if err != nil {
								log.Error(err)
								return
							}
						}

					}
				case <-quit:
					ticker.Stop()
					return
				}
			}
		}(userData)
	}

	if request.TickerStatus == "stop" {
		if ticker != nil {
			close(quit)
			wg.Wait()
		}
	}

	return nil
}

func (s *service) MessageRecipients(ctx *fiber.Ctx) ([]messagedto.MessageRecipientsResponse, error) {
	userData, err := s.userRepo.FindAll(map[string]any{"status": "sent"})
	if err != nil && err.Error() != "record not found" {
		log.Error(err)
		return nil, errors.New(i18n.CreateMsg(ctx, messages.RecordNotFound))
	}

	if err != nil {
		log.Error(err)
		return nil, errors.New(i18n.CreateMsg(ctx, messages.UnexpectedError))
	}

	var response []messagedto.MessageRecipientsResponse
	for _, user := range userData {
		response = append(response, messagedto.MessageRecipientsResponse{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Content:     user.Content,
			Status:      string(user.Status),
		})
	}

	return response, nil
}
