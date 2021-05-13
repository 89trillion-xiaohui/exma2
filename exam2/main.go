package main

import (
	"github.com/gin-gonic/gin"
	"test/exam2/calculator"
)

func main() {

	/*
	a := "123-"
	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("转换失败")
	}
	fmt.Println(b)

	 */

	// 接收数据

	r := gin.Default()
	r.GET("/calculator", func(c *gin.Context) {
		query := c.Query("key")
		calculator.Conv(query)
		res := calculator.Calculate(calculator.Conv(query))
		c.JSON(200, res)
	})
	r.Run()


/*
	str := "11+2*10+3/3*2-1"
	x := calculator.Conv(str)
	fmt.Println(calculator.Calculate(x))

 */
}
