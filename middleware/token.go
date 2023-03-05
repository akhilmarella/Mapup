package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var secretKey = []byte("nothingelse")

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Token") != "" {
			token, err := jwt.Parse(c.GetHeader("Token"), func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					log.Error().Any("action", "middleware_token.go_IsAuthorized")
					c.JSON(http.StatusUnauthorized, gin.H{"message": "error in jwt signing method"})
					return nil, fmt.Errorf("err in jwt signing")
				}
				return secretKey, nil
			})

			if err != nil {
				log.Error().Err(err).Any("token", token).Any("action", "middleware_token.go_IsAuthorized").
					Msg("error in parsing token")
				c.JSON(http.StatusUnauthorized, gin.H{"message": "error in pasrsing token"})
				c.Abort()
			}

			if token.Valid {
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					// for now iam just adding id
					// if you need anything you can add
					Id, ok := claims["id"].(string)
					if !ok {
						log.Error().Any("action", "middleware_token.go_IsAuthorized").
							Msg(" id is not found")
						c.JSON(http.StatusBadRequest, gin.H{"message": "id is not found"})
						c.Abort()
					}

					if Id == "" {
						log.Error().Any("id", Id).Any("action", "middleware_token.go_IsAuthorized").
							Msg("id is empty")
						c.JSON(http.StatusBadRequest, gin.H{"message": "id is empty"})
						c.Abort()
					}
					c.Writer.Header().Add("id", Id)
					c.Next()
				}
			} else {
				c.AbortWithError(http.StatusNotAcceptable, fmt.Errorf(" token is not valid"))
			}
		} else {
			log.Error().Any("action", "middleware_token.go_IsAuthorized").Msg("header is empty")
			c.JSON(http.StatusBadRequest, gin.H{"message": "misssing header in token"})
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no token in header"))
			return
		}
	}
}
