package controller

import (
	"test/exam2/service"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	query := c.Query("key")
	res := service.Calculate(service.Conv(query))
	c.JSON(200, res)
}
