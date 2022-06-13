package form

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID       string `json:"OrderID" binding:"required"`
	DestinationID string   `json:"DestinationID" binding:"required"`
}
