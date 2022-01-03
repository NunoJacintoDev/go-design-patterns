package pipeline

func power(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}
