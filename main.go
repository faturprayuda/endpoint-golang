package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"tes_kitalulus/controller"
	// "net/http"
)

func main() {
	r := gin.Default()
	r.GET("/list-user", controller.ListUser)
	r.GET("/detail-user/:login", controller.DetailUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}