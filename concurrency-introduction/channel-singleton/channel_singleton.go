package channel_singleton

// used to communicate with the action of adding one to the count,
// and receives a bool type just to signal "add one"
var addCh chan bool = make(chan bool)

// a channel that will receive the current value of the count
var getCountCh chan chan int = make(chan chan int)

// will communicate to the Goroutine that it should end its infinite loop
// and finish itself too
var quitCh chan bool = make(chan bool)

func init() {
	var count int
	go func(addCh <-chan bool /*getCountCh <-chan chan int,*/, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			// case ch := <-getCountCh:
			// 	ch <- count
			case <-quitCh:
				return
			}
		}
	}(addCh /*getCountCh,*/, quitCh)
}
