package controllers

import (
	"AgringBackend/config"
	"AgringBackend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOK(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func GetVehicles(c *gin.Context) {
	var vehicles []models.Vehicle
	config.DB.Find(&vehicles)
	c.JSON(http.StatusOK, vehicles)
}

func CreateVehicle(c *gin.Context) {
	var vehicle models.Vehicle
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&vehicle)
	c.JSON(http.StatusCreated, vehicle)
}

func UpdateVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	if err := config.DB.First(&vehicle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&vehicle)
	c.JSON(http.StatusOK, vehicle)
}

func DeleteVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	if err := config.DB.First(&vehicle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	config.DB.Delete(&vehicle)
	c.JSON(http.StatusNoContent, nil)
}
