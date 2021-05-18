package main

import (
	"github.com/gin-gonic/gin"
	"test/exam2/controller"
)

func main() {
	r := gin.Default()

	r.GET("/calculator", controller.Get)

	r.Run()
}
