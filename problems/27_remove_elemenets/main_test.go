package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestTwo(t *testing.T) {
	array := []int{0, 1, 2, 2, 3, 0, 4, 2}
	expectedNums := []int{0, 1, 4, 0, 3}
	k := 2
	actualLength := 5
	if err := testFunc(array, k, expectedNums, actualLength); err != nil {
		t.Error(err)
	}
}

func TestOne(t *testing.T) {
	array := []int{3, 1, 1, 3}
	expectedNums := []int{1, 1}
	k := 3
	actualLength := 2
	if err := testFunc(array, k, expectedNums, actualLength); err != nil {
		t.Error(err)
	}
}

func TestNull(t *testing.T) {
	array := []int{}
	expectedNums := []int{}
	k := 2
	actualLength := 0
	if err := testFunc(array, k, expectedNums, actualLength); err != nil {
		t.Error(err)
	}
}

func TestOneYes(t *testing.T) {
	array := []int{1}
	expectedNums := []int{1}
	k := 2
	actualLength := 1
	if err := testFunc(array, k, expectedNums, actualLength); err != nil {
		t.Error(err)
	}
}

func TestOneNo(t *testing.T) {
	array := []int{1}
	expectedNums := []int{}
	k := 1
	actualLength := 0
	if err := testFunc(array, k, expectedNums, actualLength); err != nil {
		t.Error(err)
	}
}

func TestManySame(t *testing.T) {
	array := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	expectedNums := []int{}
	k := 1
	actualLength := 0
	if err := testFunc(array, k, expectedNums, actualLength); err != nil {
		t.Error(err)
	}
}

func TestNoMatch(t *testing.T) {
	array := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	expectedNums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	k := 2
	actualLength := len(array)
	if err := testFunc(array, k, expectedNums, actualLength); err != nil {
		t.Error(err)
	}
}

func TestThree(t *testing.T) {
	array := []int{7, 87, 84, 37, 25, 9, 51, 38, 7, 47, 24, 12, 4, 9, 7}
	expectedNums := []int{87, 84, 37, 25, 9, 51, 38, 47, 24, 12, 4, 9}
	k := 7
	actualLength := len(expectedNums)
	if err := testFunc(array, k, expectedNums, actualLength); err != nil {
		t.Error(err)
	}
}

func testFunc(nums []int, val int, expectedNums []int, actualLength int) error {
	k := removeElement(nums, val) // Calls your implementation

	if k != len(expectedNums) {
		e := fmt.Errorf("Incorrect Number of integers returned: Got:%d, Need:%d", k, len(expectedNums))
		return e
	}
	sort.Ints(nums[:k]) // Sort the first k elements of nums
	sort.Ints(expectedNums)
	for i := 0; i < actualLength; i++ {
		if nums[i] != expectedNums[i] {
			e := fmt.Errorf("The array is sorted incorrectly at %d. Got: %d | Want: %d", i, nums[i], expectedNums[i])
			return e
		}
	}
	return nil
}
