package model

import "errors"

// Stack 栈
type Stack struct {
	Data []int
}

// Push 压栈
func (stack *Stack) Push(value int) {
	stack.Data = append(stack.Data, value)
	//fmt.Println("get")
}

// Pop 出栈
func (stack *Stack) Pop() (int, error) {

	// 判断栈空
	if len(stack.Data) == 0 {
		return -1, errors.New("nothing")
	}
	// 弹出栈顶数据
	num := stack.Data[len(stack.Data)-1]
	//fmt.Println(num)
	stack.Data = stack.Data[:len(stack.Data)-1]
	return num, nil

}
