package main

import (
	"test/exam2/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/calculator", controller.Get)

	r.Run()
}
