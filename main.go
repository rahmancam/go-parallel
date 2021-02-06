package main

import (
	"fmt"
)

func main() {
	numbers := GetTestData(10000 * 10000)
	result := GoSum(numbers)
	fmt.Println("Go Sum : ", result)
	result = PargoSum(numbers)
	fmt.Println("Pargo Sum : ", result)
}
