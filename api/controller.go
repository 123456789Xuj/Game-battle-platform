package api

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var (
	DB *gorm.DB
)

func ShowMenu(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "用户菜单",
	})
}

func Collectin(c *gin.Context) {
	service.Collectionservice()
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
func UserLanding(c *gin.Context) {
	service.LandingService()
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
func Home(c *gin.Context) {
	var a []models.League
	DB.Table("leagues").Order("Integral desc").Find(&a)
	c.JSON(http.StatusOK, gin.H{
		"排行信息": a,
	})
}
func AddFriends(c *gin.Context) {
	service.AddFriendservice()
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func CreationRoom(c *gin.Context) {
	service.Collectionservice()
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func JoinRoom(c *gin.Context) {
	service.JoinRoomservice()
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 在游戏房间数据库中加入用户名称和是否准备的字段
// 在创建房间时默认是否准备为否
// 其他人加入后也默认为否
// 接受前端返回信息将否改为是
// 在ready函数里进行判断，如果所有人都ready字段都是是且人数大于3人则回复ok准备开始游戏
func Ready(c *gin.Context) {
	service.Readyservice()
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func Start(c *gin.Context) {
	service.StartSerializer()
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
