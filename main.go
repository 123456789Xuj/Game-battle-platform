package main

import (
	"awesomeProject/models"
	"awesomeProject/routes"
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
	err = routes.NewRouts()
	if err != nil {
		fmt.Printf("%v", err)
	}
	_ = os.Setenv("GIN_MODE", "release")
	err2 := R.Run()
	if err != nil {
		fmt.Printf("%V/n", err2)
	}
}
