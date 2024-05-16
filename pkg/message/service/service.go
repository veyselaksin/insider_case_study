package service

import (
	"github.com/gofiber/fiber/v2/log"
	"insider_case_study/db/repository/userrepo"
	"insider_case_study/dto/messagedto"
	"time"
)

type Service interface {
	SendMessage(request messagedto.SendMessageRequest) error
}

type service struct {
	userRepo userrepo.UserRepository
}

func New(userRepo userrepo.UserRepository) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) SendMessage(request messagedto.SendMessageRequest) error {
	//data, err := s.userRepo.FindAll()
	//if err != nil {
	//	log.Error(err)
	//	return err
	//}

	ticker := time.NewTicker(2 * time.Minute)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Info("Ticker is running")
			}
		}
	}()

	return nil
}
