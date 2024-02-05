package main

import (
	"awesomeProject/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"os"
)

func main() {
	R := gin.Default()
	err := controller.InitMySQL()
	if err != nil {
		panic(err)
	}
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./public/images"))))
	R.GET("/ShowMenu", controller.ShowMenu)
	R.POST("/Collection", controller.Collectin)
	R.POST("/UserLanding", controller.UserLanding)

	_ = R.Group("/homepage", controller.UserLanding)
	{
		R.GET("", controller.Home)
		R.POST("/add", controller.AddFriends)
		R.POST("/creation", controller.CreationRoom)
		R.POST("/join", controller.JoinRoom)
		R.GET("/ready", controller.Ready)
		R.POST("/GameStart", controller.Start)

	}

	_ = os.Setenv("GIN_MODE", "release")
	err2 := R.Run()
	if err != nil {
		fmt.Printf("%V/n", err2)
	}
}
