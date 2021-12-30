package interpreter

type operationSub struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationSub) Read() int {
	return a.Left.Read() - a.Right.Read()
}
