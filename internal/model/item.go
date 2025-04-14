package model

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          string `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Status      string
	Priority    string
	Tags        []string  `gorm:"serializer:json"`
	DueDate     time.Time `gorm:"index"`
	CompletedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ItemReq struct {
	Title       string    `json:"title" `
	Description string    `json:"description" `
	Status      string    `json:"status" `
	Priority    string    `json:"priority" `
	Tags        []string  `json:"tags" `
	DueDate     time.Time `json:"due_date" `
}

func (i *Item) BeforeCreate(tx *gorm.DB) error {
	buf := make([]byte, 9)
	if _, err := rand.Read(buf); err != nil {
		return err
	}
	i.ID = base64.URLEncoding.EncodeToString(buf)
	return nil
}
