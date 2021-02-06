package main

import "math/rand"

var threshold int = 10000

// GetTestData helper
func GetTestData(n int) []float64 {
	nums := make([]float64, n)
	rand.Seed(42)
	for i := range nums {
		nums[i] = rand.Float64()
	}
	return nums
}
