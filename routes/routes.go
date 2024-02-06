package routes

import (
	"awesomeProject/api"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouts() error {
	r := gin.Default()
	r.GET("/ShowMenu", api.ShowMenu)
	r.POST("/Collection", api.Collectin)
	r.POST("/UserLanding", api.UserLanding)

	_ = r.Group("/homepage", middleware.JWT())
	{
		r.GET("", api.Home)
		r.POST("/add", api.AddFriends)
		r.POST("/creation", api.CreationRoom)
		r.POST("/join", api.JoinRoom)
		r.GET("/ready", api.Ready)
		r.POST("/GameStart", api.Start)

	}
	err := r.Run()
	if err != nil {
		return err
	}

	return err
}
