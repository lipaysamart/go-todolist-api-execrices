package model

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	Tags        []string  `gorm:"serializer:json"`
	DueDate     time.Time `gorm:"index" json:"due_date"`
	// CompletedAt time.Time     `json:"completed_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
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

func (i *Item) BeforeUpdate(tx *gorm.DB) error {
	if i.CreatedAt.IsZero() {
		i.CreatedAt = time.Now()
	}
	return nil
}
