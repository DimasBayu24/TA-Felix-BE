package form

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID            string `json:"UserID" binding:"required"`
	Status            string `json:"Status" binding:"required"`
	TotalPrice        int    `json:"TotalPrice" binding:"required"`
	TransportationID  string `json:"TransportationID" binding:"required"`
	TransportationQty int    `json:"TransportationQty" binding:"required"`
}
