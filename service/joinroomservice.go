package service

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JoinRoomservice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var room models.GameRoom
		var i string
		_ = c.BindJSON(&i)

		result := DB.Table("GameRooms").Where("MainUsername=?", i).First(&room)
		if result.Error != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "没有找到房间！"})
		}
		// 检查房间是否已满
		if room.Username != "" {
			c.JSON(http.StatusOK, gin.H{"msg": "房间已满"})
			return
		} else {
			room.Ready = false
			room.Username = i
			err := DB.Save(&room).Error
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库更新错误"})
				return
			}
		}

	}
}
