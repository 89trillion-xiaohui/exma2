package calculator

import (
	"errors"
	"strconv"
)

// 表达式求值



// Stack 栈
type Stack struct {
	data []int
}

// Push 压栈
func (stack *Stack) Push (value int) {

	stack.data = append(stack.data, value)
	//fmt.Println("get")
}

// Pop 出栈
func (stack *Stack) Pop() (int, error) {


	// 判断栈空
	if len(stack.data) == 0 {
		return -1, errors.New("nothing")
	}
	// 弹出栈顶数据
	num := stack.data[len(stack.data)-1]
	//fmt.Println(num)
	stack.data = stack.data[:len(stack.data)-1]
	return num, nil

}

// Conv 将传入字符串转换为数组保存
func Conv(str string) (num []string) {

	s := ""
	for i := range str{
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

	//fmt.Println(num)
	return num
}

// Calculate 计算结果
 func Calculate(num []string) int {
 	var stk *Stack
 	stk = &Stack{}
 	stk.data = []int{}
	var result  = 0
 	for i := range num {
 		ch := num[i]
 		//fmt.Println(ch)
 		//ch1, _ := strconv.Atoi(ch)
 		//stk.Push(ch1)
		switch ch {
		case "+", " ":
			continue
		case "-":
			num1, _ := strconv.Atoi(num[i+1])
			res := -1 * num1
			stk.Push(res)
			num[i+1] = " "
		case "*":
			num1, _ := stk.Pop()
			num2, _ := strconv.Atoi(num[i+1])
			res := num1 * num2
			stk.Push(res)
			num[i+1] = " "
		case "/":
			num1, _ := stk.Pop()
			num2, _ := strconv.Atoi(num[i+1])
			res := num1 / num2
			stk.Push(res)
			num[i+1] = " "
		default:
			num, _ := strconv.Atoi(ch)
			stk.Push(num)
		}
	}
	for i := 0; i < len(stk.data); i++{
		num := stk.data[i]
		result = result + num
	}
	return result
 }


/*
func main (){

	var stk Stack
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	fmt.Println(stk.Pop())


	str := "11+2*10+3/3*2"
	x := Conv(str)
	fmt.Println(Calculate(x))

}
*/
