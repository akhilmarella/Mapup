package handlers

import (
	"mapup/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetToken(c *gin.Context) {
	token, err := service.CreateToken()
	if err != nil {
		log.Error().Err(err).Any("token", token).Any("action", "handlers_token.go_GetToken").
			Msg("error in creating token ")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error in creating token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

