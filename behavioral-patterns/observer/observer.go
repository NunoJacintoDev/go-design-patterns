package observer

type Observer interface {
	Notify(m string)
}
