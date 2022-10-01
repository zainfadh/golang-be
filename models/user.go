package models

import "time"

// User ...
type User struct {
	ID        int       `json:"id" gorm:"column:id" example:"0"`
	Email     string    `json:"email" gorm:"column:email" example:"user@gmail.com"`
	Password  string    `json:"password" gorm:"column:password" example:"generated temporary" `
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at" `
	UpdateAt  time.Time `json:"updatedAt" gorm:"column:updated_at" `
	LastLogin time.Time `json:"lastLogin" gorm:"column:last_login"`
}

// TableName ...
func (t *User) TableName() string {
	return "public.user"
}
