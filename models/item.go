package models

import "time"

// Item ...
type Item struct {
	ID          int       `json:"id" gorm:"column:id" example:"0"`
	Name        string    `json:"name" gorm:"column:item_name"`
	Description string    `json:"description" gorm:"column:item_description"`
	Cost        float32   `json:"cost" gorm:"column:item_cost"`
	Price       float32   `json:"price" gorm:"column:item_price"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at" `
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at" `
}

// TableName ...
func (t *Item) TableName() string {
	return "public.item"
}
