package model

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID       string
	DestinationID string
}
