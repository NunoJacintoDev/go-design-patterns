package interpreter

type polishNotationStack []Interpreter

func (p *polishNotationStack) Push(s Interpreter) {
	*p = append(*p, s)
}
func (p *polishNotationStack) Pop() Interpreter {
	length := len(*p)
	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}
	return nil
}
