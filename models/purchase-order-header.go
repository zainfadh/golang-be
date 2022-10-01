package models

import "time"

// PurchaseOrderHeader ...
type PurchaseOrderHeader struct {
	ID          int       `json:"id" gorm:"column:id"`
	Date        time.Time `json:"date" gorm:"po_date"`
	Description string    `json:"description" gorm:"column:po_description"`
	Cost        float32   `json:"cost" gorm:"column:po_cost"`
	Price       float32   `json:"price" gorm:"column:po_price"`
	CreatedAt   time.Time `json:"createdAt" gorm:"created_at"`
	UpdateAt    time.Time `json:"updateAt" gorm:"update_at"`
}

// TableName ...
func (t *PurchaseOrderHeader) TableName() string {
	return "public.po_h"
}
