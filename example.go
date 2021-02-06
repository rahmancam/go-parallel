package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var threshold int = 10000

func sum(numbers []float64) float64 {
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
		left = sum(numbers[:half])
	}()
	right := sum(numbers[half:])
	wg.Wait()
	return left + right
}

func getTestData(n int) []float64 {
	nums := make([]float64, n)
	rand.Seed(42)
	for i := range nums {
		nums[i] = rand.Float64()
	}
	return nums
}

func main() {
	numbers := getTestData(10000 * 10000)
	result := sum(numbers)
	fmt.Println("sum : ", result)
}
