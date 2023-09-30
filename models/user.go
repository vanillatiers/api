package models

import (
	"time"

	"gorm.io/gorm"
)
type User struct {
	ID              string         `json:"id" gorm:"primaryKey"`
	Username        string         `json:"username"`
	Region          string         `json:"region"`
	PreferredServer string         `json:"preferredServer"`
	Rank            string         `json:"rank"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `json:"deletedAt"`
}
