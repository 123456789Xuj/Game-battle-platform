package service

import (
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Collectionservice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.NewUser
		var user1 models.NewUser
		var a models.League
		_ = c.BindJSON(&user)
		a.UserName = user.Username
		a.Integral = 0
		DB.Table("NewUsers").Where("username=?", user.Username).First(&user1)
		if user.Username == user1.Username {
			c.JSON(http.StatusOK, gin.H{"msg": "该用户名已占用！"})
			return
		} else {
			DB.Create(&user)
			DB.Create(&a)
			fmt.Printf("用户注册成功！用户名：%s，手机号码：%d\n", user.Username, user.PhoneNumber)
			c.JSON(http.StatusOK, gin.H{"msg": "用户注册成功！"})
		}
	}
}
