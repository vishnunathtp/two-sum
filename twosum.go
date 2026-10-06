package main

import "fmt"

// TwoSum finds indices of two numbers such that they add up to target.
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if idx, exists := seen[diff]; exists {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(nums, target)
	fmt.Printf("TwoSum(%v, %d) = %v\n", nums, target, result)
}
