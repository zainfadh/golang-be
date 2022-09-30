package models

import "time"

// User ...
type User struct {
	ID         int64     `json:"id" gorm:"column:id" example:"0"`
	Email      string    `json:"email" gorm:"column:email" example:"user@ottopay.id"`
	Password   string    `json:"password" gorm:"column:password" example:"generated temporary" `
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at" `
	LastUpdate time.Time `json:"lastUpdate" gorm:"column:last_update" `
	LastLogin  time.Time `json:"lastLogin" gorm:"column:last_login"`
}

// TableName ...
func (t *User) TableName() string {
	return "public.user"
}
