package main

import (
	"sync"
)

// GoSum example
func GoSum(numbers []float64) float64 {
	if len(numbers) < threshold {
		var sum float64
		for _, num := range numbers {
			sum += num
		}
		return sum
	}
	half := len(numbers) / 2
	var wg sync.WaitGroup
	wg.Add(1)
	var left float64
	go func() {
		defer wg.Done()
		left = GoSum(numbers[:half])
	}()
	right := GoSum(numbers[half:])
	wg.Wait()
	return left + right
}
