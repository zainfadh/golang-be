package models

// PurchaseOrderDetail ...
type PurchaseOrderDetail struct {
	ID                int     `json:"id" gorm:"column:id"`
	PurchaseOrderId   int     `json:"poId" gorm:"column:po_id"`
	ItemId            int     `json:"itemId" gorm:"column:po_item_id"`
	Quantity          int     `json:"qty" gorm:"po_item_qty"`
	PurchaseOrderCost float32 `json:"poCost" gorm:"column:po_item_cost"`
	PurchasePrice     float32 `json:"poPrice" gorm:"column:po_item_price"`
}

// TableName ...
func (t *PurchaseOrderDetail) TableName() string {
	return "public.po_d"
}
