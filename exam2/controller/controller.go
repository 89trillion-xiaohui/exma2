package controller

import (
	"github.com/gin-gonic/gin"
	"test/exam2/service"
)

func Get(c *gin.Context) {
	query := c.Query("key")
	res := service.Calculate(service.Conv(query))
	c.JSON(200, res)
}
