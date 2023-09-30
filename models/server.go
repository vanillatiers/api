package models

import (
  "time"

  "gorm.io/gorm"
)

type Server struct {
	ID                   string         `json:"id" gorm:"primaryKey"`
	Cooldown             uint           `json:"cooldown"`
	LT3PlusCooldown      uint           `json:"lt3PlusCooldown"`
	TicketAutoclose      uint           `json:"ticketAutoclose"`
	QueueAutoclose       uint           `json:"queueAutoclose"`
	QueueChannelID       string         `json:"queueChannelID"`
	LogsChannelID        string         `json:"logsChannelID"`
	EURoleID             string         `json:"euRoleID"`
	NARoleID             string         `json:"naRoleID"`
	HT1RoleID            string         `json:"ht1RoleID"`
	LT1RoleID            string         `json:"lt1RoleID"`
	HT2RoleID            string         `json:"ht2RoleID"`
	LT2RoleID            string         `json:"lt2RoleID"`
	HT3RoleID            string         `json:"ht3RoleID"`
	LT3RoleID            string         `json:"lt3RoleID"`
	HT4RoleID            string         `json:"ht4RoleID"`
	LT4RoleID            string         `json:"lt4RoleID"`
	HT5RoleID            string         `json:"ht5RoleID"`
	LT5RoleID            string         `json:"lt5RoleID"`
	EUTicketsCategoryID  string         `json:"euTicketsCategoryID"`
	NATicketsCategoryID  string         `json:"naTicketsCategoryID"`
	HT3TicketsCategoryID string         `json:"ht3TicketsCategoryID"`
	UpdatedAt            time.Time      `json:"updatedAt"`
	DeletedAt            gorm.DeletedAt `json:"deletedAt"`
}
