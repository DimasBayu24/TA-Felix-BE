package model

import (
	"gorm.io/gorm"
)

type ProductPackage struct {
	gorm.Model
	Package          string
	DestinationPlace string
	Price            int
}
