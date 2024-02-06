package service

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Creationservice() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		strUsername, ok := username.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}
		var p1 = models.GameRoom{}
		p1.Ready = false
		p1.MainUsername = strUsername
		DB.Create(&models.GameRoom{})
		c.JSON(http.StatusOK, gin.H{
			"username": strUsername,
		})
	}
}
