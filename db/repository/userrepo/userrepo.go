package userrepo

import (
	"gorm.io/gorm"
	"insider_case_study/db/models"
)

type UserRepository interface {
	FindAll() (*models.User, error)
}

type dataAccess struct {
	db *gorm.DB
	//redisDb *redis.Client
}

func New(db *gorm.DB) UserRepository {
	return &dataAccess{
		db: db,
	}
}

func (d *dataAccess) FindAll() (user *models.User, err error) {
	return user, nil
}
