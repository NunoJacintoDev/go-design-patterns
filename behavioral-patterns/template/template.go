package template

type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}
