package models

type APIKey struct {
	Key   string `gorm:"primaryKey"`
	Perms uint8
}
