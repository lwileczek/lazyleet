package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func removeElement(nums []int, val int) int {
	p := 0
	w := 0
	n := len(nums)
	for (p + w) < n {
		if nums[p] == val {
			nums[p] = nums[n-w-1]
			w++
		} else {
			p++
		}
	}
	return n - w
}
