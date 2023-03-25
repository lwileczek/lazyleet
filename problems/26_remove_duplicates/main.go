package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func removeDuplicates(nums []int) int {
	maxLen := len(nums)
	k := len(nums)
	i := 0
	for i < k {
		if i+1 >= maxLen {
			break
		}
		if nums[i] == nums[i+1] {
			k--
			for j := i; j < k; j++ {
				nums[j] = nums[j+1]
			}
		} else {
			i++
		}
	}
	if k+1 < maxLen {
		return k + 1
	}
	return maxLen
}

//J only writes when we want it to, we only do one loop through
// will automatically write at position 1 and then jump to the next slot
// wen a new value is found. Don't have to shift all the elements down.
// Just write the ones we care about.
func preferredRemoveDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

// Only have to count the writes. Even if you write every element, len is writes+1
func shortRemoveDuplicates(nums []int) int {
	count := 0
	for i := range nums {
		if nums[count] != nums[i] {
			count++
			nums[count] = nums[i]
		}
	}
	return count + 1
}
