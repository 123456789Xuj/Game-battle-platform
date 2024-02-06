package service

import (
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var board [15][15]int
var player int

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

func StartSerializer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var chess models.Chess
		//用数组表示坐标初始化一个棋盘
		borad()
		defer borad()
		player = 1
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
		var league models.League

		var b = league.Integral
		// 更新当前玩家在数据库中的准备状态
		err := DB.Table("GameRooms").Where("username=?", strUsername).Update("Ready", true).First(&ready).Error
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
					league.UserName = strUsername
					b = b + 1
					DB.Table("Leagues").Where("UserName = ?", strUsername).Update("Integral", b)
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
					DB.Table("GameRooms").Where("MainUsername=?", strUsername).First(&game)
					DB.Table("Leagues").Where("UserName = ?", game.Username).Update("Integral", b)

					break
				}
			}

		}
		//判断输赢

		//重置棋盘
	}

}
