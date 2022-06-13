package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID            string
	Status            string
	TransportationID  string
	TransportationQty int
	TotalPrice        int
}
