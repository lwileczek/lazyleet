package main

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	array := []int{1, 3, 5, 6}
	target := 5
	idx := 2
	if err := testFunc(array, target, idx); err != nil {
		t.Error(err)
	}
}

func TestTwo(t *testing.T) {
	array := []int{1, 3, 5, 6}
	target := 2
	idx := 1
	if err := testFunc(array, target, idx); err != nil {
		t.Error(err)
	}
}

func TestThree(t *testing.T) {
	array := []int{1, 3, 5, 6}
	target := 7
	idx := 4
	if err := testFunc(array, target, idx); err != nil {
		t.Error(err)
	}
}

func testFunc(nums []int, target int, actualValue int) error {
	k := searchInsert(nums, target) // Calls your implementation

	if k != actualValue {
		e := fmt.Errorf("Error! Expected %d but got %d. Input: %v | %d", actualValue, k, nums, target)
		return e
	}
	return nil
}

func BenchmarkSearch(b *testing.B) {
	inputArray := make([]int, 1000)
	for z := 0; z < 1000; z++ {
		inputArray[z] = z
	}
	for i := 0; i < b.N; i++ {
		searchInsert(inputArray, 987)
	}
}
