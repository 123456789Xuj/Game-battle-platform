package controller

import (
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var (
	DB     *gorm.DB
	R      gin.Engine
	roomId int
	board  [15][15]int
	player int
)

func ShowMenu(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "用户菜单",
	})
}
func InitMySQL() error {
	dns := "root:xjl666nbsg@tcp(127.0.0.1:3306)/landing?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = DB.AutoMigrate(&models.NewUser{})
	_ = DB.AutoMigrate(&models.GameRoom{})
	_ = DB.AutoMigrate(&models.GameRoom{})
	_ = DB.AutoMigrate(&models.Chess{})
	_ = DB.AutoMigrate(&models.League{})
	return nil

}

func Collectin(c *gin.Context) {
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
func UserLanding(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var user models.NewUser
	result := DB.Table("NewUsers").Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"meg": "登录成功！"},
		)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "登录失败，请检查用户名和密码是否正确。"})
	}

	c.SetCookie("sessionID", username, 0, "/", "", false, true)
	c.Next()
}
func Home(c *gin.Context) {
	var a []models.League
	DB.Table("leagues").Order("Integral desc").Find(&a)
	c.JSON(http.StatusOK, gin.H{
		"排行信息": a,
	})
}
func AddFriends(c *gin.Context) {
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

func CreationRoom(c *gin.Context) {
	sessionID, err := c.Cookie("sessionID")
	if err != nil {
		// 未找到有效的会话ID，处理错误
		c.JSON(http.StatusOK, gin.H{"msg": "未找到有效id"})
	}
	var p1 = models.GameRoom{}
	p1.Ready = false
	p1.MainUsername = sessionID
	DB.Create(&models.GameRoom{})
	c.JSON(http.StatusOK, gin.H{
		"username": sessionID,
	})
}

func JoinRoom(c *gin.Context) {
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

// 在游戏房间数据库中加入用户名称和是否准备的字段
// 在创建房间时默认是否准备为否
// 其他人加入后也默认为否
// 接受前端返回信息将否改为是
// 在ready函数里进行判断，如果所有人都ready字段都是是且人数大于3人则回复ok准备开始游戏
func Ready(c *gin.Context) {
	// 获取当前玩家的用户名
	sessionID, err := c.Cookie("sessionID")
	if err != nil {
		// 未找到有效的会话ID，处理错误
		c.JSON(http.StatusOK, gin.H{"msg": "未找到有效id"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"username": sessionID})
	}
	var ready models.GameRoom
	// 更新当前玩家在数据库中的准备状态
	err = DB.Table("GameRooms").Where("username=?", sessionID).Update("Ready", true).First(&ready).Error
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

func borad() {
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			board[i][j] = 0
		}
	}

}
func drop(x, y int) {
	board[x][y] = 1
	player = 3 - player
}
func chickwin(x, y int) bool {
	//判断横向是否连成五子
	var (
		b   int
		win bool
	)
	if x >= 4 && x < 11 {
		for i := 0; i < 5; i++ {
			for j := -4; j < 1; j++ {
				a := board[x+i+j][y]
				b = a + b
				if b == 5 {
					win = true
					return win
				}
			}
		}

	}

	//判断纵向是否连成五子
	if y >= 4 && y < 11 {
		for i := 0; i < 5; i++ {
			for j := -4; j < 1; j++ {
				a := board[x][y+i+j]
				b = a + b
				if b == 5 {
					win = true
					return win
				}
			}
		}

	}
	//判断从左上到右下是否连成五子
	if y >= 4 && x >= 4 && y < 11 && x < 11 {
		for i := 0; i < 5; i++ {
			for j := -4; j < 1; j++ {
				a := board[x+i+j][y+i+j]
				b = a + b
				if b == 5 {
					win = true
					return win
				}
			}
		}

	}
	//判断从右上到坐下是否连成五子
	if y >= 4 && x >= 4 && y < 11 && x < 11 {
		for i := 0; i < 5; i++ {
			for j := -4; j < 1; j++ {
				a := board[x+i+j][y-i-j]
				b = a + b
				if b == 5 {
					win = true
					return win
				}
			}
		}

	}
	return false
}
func Start(c *gin.Context) {
	var chess models.Chess
	//用数组表示坐标初始化一个棋盘
	borad()
	defer borad()
	player = 1
	sessionID, err := c.Cookie("sessionID")
	if err != nil {
		// 未找到有效的会话ID，处理错误
		c.JSON(http.StatusOK, gin.H{"msg": "未找到有效id"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"username": sessionID})
	}
	var ready models.GameRoom
	var league models.League

	var b = league.Integral
	// 更新当前玩家在数据库中的准备状态
	err = DB.Table("GameRooms").Where("username=?", sessionID).Update("Ready", true).First(&ready).Error
	if err != nil {
		// 处理数据库更新错误
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库更新错误"})
		return
	}
	//落子：接收前端返回的坐标信息   交换棋手
	for {
		err := c.BindJSON(chess)
		if err != nil {
			fmt.Printf("%v/d", err)
		}
		x := chess.X
		y := chess.Y
		if player == 1 {
			c.JSON(http.StatusOK, gin.H{"执棋手": "黑棋"})
			drop(x, y)
			chickwin(x, y)
			if chickwin(x, y) == true {
				league.UserName = sessionID
				b = b + 1
				DB.Table("Leagues").Where("UserName = ?", sessionID).Update("Integral", b)
				break
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"执棋手": "白棋"})
			drop(x, y)
			chickwin(x, y)
			if chickwin(x, y) == true {
				var game models.GameRoom
				league.UserName = game.Username
				b = b + 1
				DB.Table("GameRooms").Where("MainUsername=?", sessionID).First(&game)
				DB.Table("Leagues").Where("UserName = ?", game.Username).Update("Integral", b)

				break
			}
		}

	}
	//判断输赢

	//重置棋盘

}
