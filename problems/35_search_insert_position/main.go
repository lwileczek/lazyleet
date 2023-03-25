package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func searchInsert(nums []int, target int) int {
	sliceLength := len(nums)
	for z := 0; z < sliceLength; z++ {
		if target <= nums[z] {
			return z
		}
	}
	return sliceLength
}

//fastSearchInsert, checking endpoints and moving in is the same amount
//of checks but less loops. If the answer is towards the end it will find
//it more sooner.
func fastSearchInsert(nums []int, target int) int {
	nl := len(nums)
	left := 0
	right := nl - 1
	if nl == 1 {
		if target > nums[0] {
			return 1
		}
		return 0
	}
	for true {
		if target <= nums[left] {
			return left
		} else if target > nums[right] {
			return right + 1
		} else {
			left += 1
			right -= 1
		}
	}
	return 0
}
