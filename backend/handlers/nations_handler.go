package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/model"
)

func HandleNations(c *gin.Context) {
	c.JSON(http.StatusOK, model.GetNations())
}
