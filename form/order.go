package form

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Status               string     `json:"Status" binding:"required"`
	TotalPrice           int        `json:"TotalPrice" binding:"required"`
	TransportationID     int        `json:"TransportationID" binding:"required"`
	TransportationQty    int        `json:"TransportationQty" binding:"required"`
	Fullname             string     `json:"Fullname"`
	DestinationPackageID *int       `json:"DestinationPackageID"`
	IsPackage            bool       `json:"IsPackage"`
	Email                string     `json:"Email"`
	Phone                string     `json:"Phone"`
	OrderDate            *time.Time `json:"OrderDate"`
	Duration             int        `json:"Duration"`
	PictureUrl           string     `json:"PictureUrl"`
}
