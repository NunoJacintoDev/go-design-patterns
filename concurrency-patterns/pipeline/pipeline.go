package pipeline

func LaunchPipeline(amount int) int {
	firstCh := generator(amount)
	secondCh := power(firstCh)
	thirdCh := sum(secondCh)

	result := <-thirdCh
	return result
}
