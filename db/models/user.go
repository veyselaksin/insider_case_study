package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	PhoneNumber string    `gorm:"type:varchar(20);unique;not null"`
	Content     string    `gorm:"type:varchar(100);not null"`
	StatusID    uuid.UUID `gorm:"type:uuid;not null"`
	Abstract

	// Relationships
	Status Status `gorm:"foreignKey:StatusID"`
}

func (u *User) TableName() string {
	return "public.users"
}
