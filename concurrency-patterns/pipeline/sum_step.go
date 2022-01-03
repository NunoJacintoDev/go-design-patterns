package pipeline

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}
