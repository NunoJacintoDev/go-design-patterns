package log_appender_example

type Visitable interface {
	Accept(Visitor)
}
