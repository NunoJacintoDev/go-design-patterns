package pipeline

func generator(max int) <-chan int {
	outChInt := make(chan int, 100)
	go func() {
		for i := 1; i <= max; i++ {
			outChInt <- i
		}
		close(outChInt)
	}()
	return outChInt
}
