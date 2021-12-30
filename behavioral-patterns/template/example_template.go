package template

import "strings"

type ExampleTemplate struct{}

func (t *ExampleTemplate) first() string {
	return "hello"
}
func (t *ExampleTemplate) third() string {
	return "template"
}

func (t *ExampleTemplate) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}
