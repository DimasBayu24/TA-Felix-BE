package form

import "gorm.io/gorm"

type ProductPackage struct {
	gorm.Model
	Package          string `json:"Package" binding:"required"`
	DestinationPlace string `json:"DestinationPlace" binding:"required"`
	Price            int    `json:"Price" binding:"required"`
}
