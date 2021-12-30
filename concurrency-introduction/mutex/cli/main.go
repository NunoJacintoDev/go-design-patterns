package main

/*
The concurrent counter


*/
import (
	"go-design-patterns/concurrency-introduction/mutex"
	"time"
)

func main() {
	counter := mutex.Counter{}
	for i := 0; i < 10; i++ {
		go func() {
			counter.Lock() // <- comment this and run "make mutex"
			counter.Value++
			defer counter.Unlock() // <- comment this and run "make mutex"
		}()
	}
	time.Sleep(time.Second)
	counter.Lock()
	defer counter.Unlock()
	println(counter.Value)
}
