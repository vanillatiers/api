package models

import (
	"time"
)

type Queue struct {
	MessageID string    `json:"messageID" gorm:"primaryKey"`
	Members   []string  `json:"members" gorm:"serializer:json"`
	Testers   []string  `json:"testers" gorm:"serializer:json"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
