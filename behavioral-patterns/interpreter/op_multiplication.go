package interpreter

type operationMul struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationMul) Read() int {
	return a.Left.Read() * a.Right.Read()
}
