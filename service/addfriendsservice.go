package service

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddFriendservice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var username string
		//按照username查找好友（在数据库中按username查找）
		_ = c.BindJSON(&username)
		var user models.NewUser
		DB.Table("NewUser").Where("username=?", username).First(&user)
		//添加好友
		var GameFriend models.GameFriend
		GameFriend.Username = user.Username
		DB.Create(&GameFriend)
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	}
}
