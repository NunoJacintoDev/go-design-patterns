package interfaces

type Visitable interface {
	Accept(Visitor)
}
