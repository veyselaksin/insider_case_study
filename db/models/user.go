package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"insider_case_study/helpers/enums"
)

type User struct {
	ID          string       `gorm:"type:uuid;primary_key"`
	PhoneNumber string       `gorm:"type:varchar(20);unique;not null"`
	Content     string       `gorm:"type:varchar(100);not null"`
	Status      enums.Status `gorm:"type:varchar(20);not null"`
	Abstract
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}

func (u *User) TableName() string {
	return "public.users"
}
