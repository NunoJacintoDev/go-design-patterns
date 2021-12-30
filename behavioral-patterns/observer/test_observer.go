package observer

import "fmt"

type TestObserver struct {
	ID      int
	Message string
}

func (p *TestObserver) Notify(m string) {
	fmt.Printf("Observer %d: message '%s' received \n", p.ID, m)
	p.Message = m
}
