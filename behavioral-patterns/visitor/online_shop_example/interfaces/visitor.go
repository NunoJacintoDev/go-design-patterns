package interfaces

type Visitor interface {
	Visit(Retriever)
}
