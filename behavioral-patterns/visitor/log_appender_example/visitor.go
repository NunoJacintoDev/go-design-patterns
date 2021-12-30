package log_appender_example

type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}
