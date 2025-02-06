package handlers

import (
	"AgringBackend/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFieldWorkersHandler(c *gin.Context) {
	workers, err := repositories.GetVirtualFieldWorkers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, workers)
}
