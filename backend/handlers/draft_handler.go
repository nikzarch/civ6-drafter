package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/services"
	"strconv"
)

func HandleDrafts(c *gin.Context) {
	players_amount := c.Query("players")
	choicesForPlayer := c.Query("choices")
	if players_amount == "" || choicesForPlayer == "" {
		c.JSON(http.StatusBadRequest, "Players or/and choices not found")
		return
	}
	playersAmountInt, err := strconv.Atoi(players_amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	choicesForPlayerInt, err := strconv.Atoi(choicesForPlayer)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	draft, err := services.GetDraft(playersAmountInt, choicesForPlayerInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, draft)
}
