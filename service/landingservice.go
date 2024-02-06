package service

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

var DB *gorm.DB

func LandingService() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		var user models.NewUser
		result := DB.Table("NewUsers").Where("username = ? AND password = ?", username, password).First(&user)
		if result.Error == nil {
			expiresAt := time.Now().Add(time.Hour * 24).Unix()
			claims := jwt.MapClaims{
				"username": username,
				"exp":      expiresAt,
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString([]byte("secret_key"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}
			go c.Set("username", claims["username"].(string))
			c.JSON(http.StatusOK, gin.H{"token": tokenString})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 2002,
				"msg":  "鉴权失败",
			})
			return
		}
	}
}
