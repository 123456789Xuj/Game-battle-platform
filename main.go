package main

import (
	"awesomeProject/api"
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func main() {
	R := gin.Default()
	err := models.InitMySQL()
	if err != nil {
		panic(err)
	}
	R.GET("/ShowMenu", api.ShowMenu)
	R.POST("/Collection", api.Collectin)
	R.POST("/UserLanding", api.UserLanding)

	_ = R.Group("/homepage", api.UserLanding)
	{
		R.GET("", api.Home)
		R.POST("/add", api.AddFriends)
		R.POST("/creation", api.CreationRoom)
		R.POST("/join", api.JoinRoom)
		R.GET("/ready", api.Ready)
		R.POST("/GameStart", api.Start)

	}

	_ = os.Setenv("GIN_MODE", "release")
	err2 := R.Run()
	if err != nil {
		fmt.Printf("%V/n", err2)
	}
}
