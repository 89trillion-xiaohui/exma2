package calculator

import (
	"fmt"
	"testing"
)

func TestConv(t *testing.T) {
	var str = "1+2*10+3/3-5"
	i := Conv(str)
	fmt.Println(i)
}

func TestCalculate(t *testing.T) {
	var str = "1+2*10+3/3-5"
	i := Conv(str)
	j := Calculate(i)
	fmt.Println(j)
}