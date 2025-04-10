package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lipaysamart/go-todolist-api-execrices/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserLoginReq struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type UserRegisterReq struct {
	Email    string `json:"email" `
	Username string `json:"username" `
	Password string `json:"password" `
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	u.Password = utils.HashAndSalt([]byte(u.Password))
	return nil
}
