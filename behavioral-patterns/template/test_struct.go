package template

type TestStruct struct {
	ExampleTemplate
}

func (m *TestStruct) Message() string {
	return "world"
}
