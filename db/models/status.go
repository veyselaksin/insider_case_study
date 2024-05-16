package models

import "github.com/google/uuid"

type Status struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key"`
	Name string    `gorm:"type:varchar(20);not null"`
	Abstract
}

func (s *Status) TableName() string {
	return "public.statuses"
}
