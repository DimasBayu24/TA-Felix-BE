package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-api/db"
	"product-api/form"
	"product-api/model"
)

func GetAllDestinations(c *gin.Context) {
	var product []model.Destination
	db.DB.Find(&product)

	c.JSON(http.StatusOK, product)
}

func GetDestinationByID(c *gin.Context) {
	var product model.Destination
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func PostDestination(c *gin.Context) {
	var input form.Destination
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product := model.Destination{
		Place:       input.Place,
		PlaceOption: input.PlaceOption,
		Price:       input.Price,
	}
	db.DB.Create(&product)

	c.JSON(http.StatusCreated, product)
}

func UpdateDestinationByID(c *gin.Context) {
	var product model.Destination
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.Destination
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



	db.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, product)
}

func DeleteDestinationByID(c *gin.Context) {
	var product model.Destination
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
