package models

import (
	"time"

	"gorm.io/gorm"
)

type Tester struct {
	UserID     string         `json:"id" gorm:"primaryKey"`
	TotalTests uint           `json:"testsInTotal"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt"`
}
