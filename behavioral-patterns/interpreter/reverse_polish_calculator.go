package interpreter

import (
	"strconv"
	"strings"
)

// Calculate Facade for Reverse Polish operation
func Calculate(o string) (int, error) {
	stack := polishNotationStack{}
	operators := strings.Split(o, " ")
	for _, operatorString := range operators {
		if isOperation(operatorString) {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := operatorFactory(operatorString, left, right)
			res := value(mathFunc.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				panic(err)
			}
			temp := value(val)
			stack.Push(&temp)
		}
	}
	return int(stack.Pop().Read()), nil
}

func isOperation(op string) bool {
	return op == SUM ||
		op == SUB ||
		op == MUL ||
		op == DIV
}
