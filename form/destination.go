package form

import "gorm.io/gorm"

type Destination struct {
	gorm.Model
	Place       string `json:"Place" binding:"required"`
	PlaceOption string `json:"PlaceOption" binding:"required"`
	Price       int `json:"Price" binding:"required"`
}
