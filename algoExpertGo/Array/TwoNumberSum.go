package main

import "fmt"

func TwoNumberSum(array []int, target int) []int {
	seen := map[int]bool{}
	for _, val := range array {
		targetVal := target - val
		if _, ok := seen[targetVal]; ok {
			return []int{val, targetVal}
		}
		seen[val] = true
	}
	return []int{}
}

func TestTwoNumbereSum() {
	array := [...]int{3, 5, -4, 8, 11, 1, -1, 6}
	fmt.Println(TwoNumberSum(array[:], 10))
}
