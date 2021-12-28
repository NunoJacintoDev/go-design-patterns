package bridge

type PrinterAPI interface {
	PrintMessage(string) error
}
