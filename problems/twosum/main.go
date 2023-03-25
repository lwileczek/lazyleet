package main

import "fmt"

func main() {
	n := []int{2, 7, 15, 11}
	idx := twoSum(n, 9)
	fmt.Println(idx)

	//Two
	nums := []int{3, 2, 4}
	idx1 := twoSum(nums, 6)
	fmt.Println(idx1)

	//Three
	numbers := []int{3, 3}
	idx2 := twoSum(numbers, 6)
	fmt.Println(idx2)

	//Test four
	test4 := []int{0, 4, 3, 0}
	ans4 := twoSum(test4, 0)
	fmt.Println(ans4)
}

// This is best answer. once through the list, no sort, accounts for negative numbers, O(n) complexity
// Lookups are O(1) unless their is a collision (unlikley)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, x := range nums {
		compliment := target - x
		if idx, ok := m[compliment]; ok {
			return []int{idx, i}
		}
		m[x] = i
	}
	// it will return before here.
	return []int{}
}
