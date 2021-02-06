package main

import (
	"github.com/exascience/pargo/parallel"
)

// PargoSum example
func PargoSum(numbers []float64) float64 {
	return parallel.RangeReduceFloat64(0, len(numbers), 0, func(low, high int) float64 {
		var sum float64
		for i := low; i < high; i++ {
			sum += numbers[i]
		}
		return sum
	}, func(left, right float64) float64 {
		return left + right
	})
}
