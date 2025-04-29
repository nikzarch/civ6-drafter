package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/model"
)

func HandleNations(c *gin.Context) {
	nations, err := model.GetNations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nations)
}
