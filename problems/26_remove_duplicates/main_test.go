package main

import "testing"

func TestOneTwo(t *testing.T) {
	array := []int{1, 1, 2}
	expectedNums := []int{1, 2}
	k := removeDuplicates(array)
	if k != len(expectedNums) {
		t.Fatalf("Incorrect Number of integers returned: Got:%d, Need:%d", k, len(expectedNums))
	}
	for i := 0; i < k; i++ {
		if array[i] != expectedNums[i] {
			t.Fatalf("The array is sorted incorrectly at %d. Got: %d | Want: %d", i, array[i], expectedNums[i])
		}
	}
}

func TestExampleII(t *testing.T) {
	array := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedNums := []int{0, 1, 2, 3, 4}
	k := removeDuplicates(array)
	if k != len(expectedNums) {
		t.Fatalf("Incorrect Number of integers returned: Got:%d, Need:%d", k, len(expectedNums))
	}
	for i := 0; i < k; i++ {
		if array[i] != expectedNums[i] {
			t.Fatalf("The array is sorted incorrectly at %d. Got: %d | Want: %d", i, array[i], expectedNums[i])
		}
	}
}

func TestThreePeat(t *testing.T) {
	array := []int{0, 0, 0, 1, 1, 1, 2, 2}
	expectedNums := []int{0, 1, 2}
	k := removeDuplicates(array)
	if k != len(expectedNums) {
		t.Fatalf("Incorrect Number of integers returned: Got:%d, Need:%d", k, len(expectedNums))
	}
	for i := 0; i < k; i++ {
		if array[i] != expectedNums[i] {
			t.Fatalf("The array is sorted incorrectly at %d. Got: %d | Want: %d", i, array[i], expectedNums[i])
		}
	}
}

func TestAllSame(t *testing.T) {
	array := []int{0, 0, 0, 0, 0, 0, 0, 0}
	expectedNums := []int{0}
	k := removeDuplicates(array)
	if k != len(expectedNums) {
		t.Fatalf("Incorrect Number of integers returned: Got:%d, Need:%d", k, len(expectedNums))
	}
	for i := 0; i < k; i++ {
		if array[i] != expectedNums[i] {
			t.Fatalf("The array is sorted incorrectly at %d. Got: %d | Want: %d", i, array[i], expectedNums[i])
		}
	}
}

func TestAllGood(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 6, 8, 9}
	expectedNums := []int{0, 1, 2, 3, 4, 6, 8, 9}
	k := removeDuplicates(array)
	if k != len(expectedNums) {
		t.Fatalf("Incorrect Number of integers returned: Got:%d, Need:%d", k, len(expectedNums))
	}
	for i := 0; i < k; i++ {
		if array[i] != expectedNums[i] {
			t.Fatalf("The array is sorted incorrectly at %d. Got: %d | Want: %d", i, array[i], expectedNums[i])
		}
	}
}
