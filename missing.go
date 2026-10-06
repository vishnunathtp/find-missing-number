package main

import "fmt"

func MissingNumber(nums []int) int {
	n := len(nums)
	xorSum := n
	for i, num := range nums {
		xorSum ^= i ^ num
	}
	return xorSum
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Printf("Missing number in %v is: %d\n", nums, MissingNumber(nums))
}
