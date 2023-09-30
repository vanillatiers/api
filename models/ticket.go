package models

import "time"

type Ticket struct {
	ChannelID string    `json:"channelID" gorm:"primaryKey"`
	UserID    string    `json:"memberID"`
	Testers   []string  `json:"testers" gorm:"serializer:json"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
