package service

import (
	"fmt"
	"strconv"

	model2 "test/exam2/model"
)

// Conv 将传入字符串转换为数组保存
func Conv(str string) (num []string) {
	s := ""
	for i := range str {
		ch := string(str[i])
		switch ch {
		case " ":
			continue
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			s += ch
		case "*", "/", "+", "-":
			num = append(num, s)
			s = ""
			num = append(num, ch)
		}
	}
	num = append(num, s)
	return num
}

// Calculate 计算结果
func Calculate(num []string) int {
	var stk *model2.Stack
	stk = &model2.Stack{}
	stk.Data = []int{}
	var result = 0
	for i := range num {
		ch := num[i]
		switch ch {
		case " ":
			continue
		case "+":

		case "-":
			num1, err := strconv.Atoi(num[i+1])
			if err != nil {
				fmt.Println("表达式错误")
				return 0
			}
			res := -1 * num1
			stk.Push(res)
			num[i+1] = " "
		case "*":
			num1, _ := stk.Pop()
			num2, err := strconv.Atoi(num[i+1])
			if err != nil {
				fmt.Println("表达式错误")
				return 0
			}
			res := num1 * num2
			stk.Push(res)
			num[i+1] = " "
		case "/":
			num1, _ := stk.Pop()
			num2, err := strconv.Atoi(num[i+1])
			if num2 == 0 {
				fmt.Println("表达式错误")
				return 0
			}
			if err != nil {
				fmt.Println("表达式错误")
				return 0
			}
			res := num1 / num2
			stk.Push(res)
			num[i+1] = " "
		default:
			num, _ := strconv.Atoi(ch)
			stk.Push(num)
		}
	}
	for i := 0; i < len(stk.Data); i++ {
		num := stk.Data[i]
		result = result + num
	}
	return result
}
