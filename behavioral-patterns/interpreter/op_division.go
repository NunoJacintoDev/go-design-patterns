package interpreter

type operationDiv struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationDiv) Read() int {
	return a.Left.Read() / a.Right.Read()
}
