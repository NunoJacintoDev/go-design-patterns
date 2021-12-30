/*
A TV channel in real life is something that
connects an emission (from a studio) to millions of TVs (the receivers).

Channels in Go work in a similar fashion.

One or more Goroutines can work as emitters, and one or more Goroutine can act as receivers.

One more thing channels, by default, block the execution of Goroutines until something is received. It is as if our favourite TV show delays the emission until we turn the TV on so we don't miss anything.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan string)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		channel <- "Hello World!"
		println("Finishing goroutine")
		waitGroup.Done()
	}()
	time.Sleep(time.Second)
	message := <-channel
	fmt.Println(message)
	waitGroup.Wait()
}
