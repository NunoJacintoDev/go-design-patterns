package interpreter

func operatorFactory(o string, left, right Interpreter) Interpreter {
	switch o {
	case SUM:
		return &operationSum{
			Left:  left,
			Right: right,
		}
	case SUB:
		return &operationSub{
			Left:  left,
			Right: right,
		}
	case MUL:
		return &operationMul{
			Left:  left,
			Right: right,
		}
	case DIV:
		return &operationDiv{
			Left:  left,
			Right: right,
		}
	}

	return nil
}
