package service

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取当前玩家的用户名

func Readyservice() gin.HandlerFunc {
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
		var ready models.GameRoom
		// 更新当前玩家在数据库中的准备状态
		err := DB.Table("GameRooms").Where("username=?", strUsername).Update("Ready", true).First(&ready).Error
		if err != nil {
			// 处理数据库更新错误
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库更新错误"})
			return
		}

		if ready.Username == "" {
			// 还有未准备好的玩家，或者人数不足2人，等待其他玩家准备
			c.JSON(http.StatusOK, gin.H{"msg": "房间人数未满"})
		} else {
			// 所有玩家已准备就绪，重定向到开始游戏的接口
			c.Redirect(http.StatusFound, "/Start")
		}
	}
}
