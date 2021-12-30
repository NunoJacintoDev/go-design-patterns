package interpreter

type value int

func (v *value) Read() int {
	return int(*v)
}
