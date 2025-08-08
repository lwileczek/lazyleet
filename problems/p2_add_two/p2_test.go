package main

import (
	"testing"
)

func TestToLinkedList(t *testing.T) {
	type TestCase struct {
		input    []int
		expected string
	}

	testCases := []TestCase{
		{
			input: []int{2, 4, 3},
			expected: "342",
		},
		{
			input: []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
			expected: "1000000000000000000000000000001",
		},
	}

	for _, tc := range testCases {
		linkedList := FromSlice(tc.input)
		i := linkedList.NumString()
		if i != tc.expected {
			t.Errorf("failed to construct linked list\n\tExpected %s, got %s\n\t%+v", tc.expected, i, linkedList)
		}
	}
}

func TestFromLinkedList(t *testing.T) {
	type TestCase struct {
		List     *ListNode
		Expected int
		Msg      string
	}

	testCases := []TestCase{
		{
			List: &ListNode{
				Val: 2,
			},
			Expected: 2,
			Msg:      "singleton",
		},
		{
			List: &ListNode{ },
			Expected: 0,
			Msg:      "nil case",
		},
		{
			List: FromSlice([]int{2, 4, 3}),
			Expected: 342,
			Msg:      "single slice",
		},
	}




	for _, tc := range testCases {
		i := ToInt(tc.List)
		if i != tc.Expected {
			t.Errorf("Failed to convert list to an int\n\tgot %d but expected %d\n", i, tc.Expected)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	type TestCase struct {
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}

	testCases := []TestCase{
		{
			l1:       &ListNode{},
			l2:       &ListNode{},
			expected: &ListNode{Val: 0},
		},
		{
			l1:       FromSlice([]int{2, 4, 3}),
			l2:       FromSlice([]int{5, 6, 4}),
			expected: FromSlice([]int{7, 0, 8}),
		},
		{
			l1:       FromSlice([]int{9,9,9,9,9,9,9}),
			l2:       FromSlice([]int{9,9,9,9}),
			expected: FromSlice([]int{8,9,9,9,0,0,0,1}),
		},
		{
			l1:       FromSlice([]int{1}),
			l2:       FromSlice([]int{1}),
			expected: FromSlice([]int{2}),
		},
		{
			l1:       FromSlice([]int{5}),
			l2:       FromSlice([]int{5}),
			expected: FromSlice([]int{0, 1}),
		},
		{
			l1:       FromSlice([]int{3, 2, 1}),
			l2:       FromSlice([]int{5, 4}),
			expected: FromSlice([]int{8, 6, 1}),
		},
		{
			l1:       FromSlice([]int{9, 9}),
			l2:       FromSlice([]int{1}),
			expected: FromSlice([]int{0, 0, 1}),
		},
		{
			l1:       FromSlice([]int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}),
			l2:       FromSlice([]int{5,6,4}),
			expected: FromSlice([]int{6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}),
		},
	}

	for _, tc := range testCases {
		ans := addTwoNumbers(tc.l1, tc.l2)
		if ans.String() != tc.expected.String() {
			t.Errorf("Failed to add two numbers\n\tgot %s but expected %s\n", ans.String(), tc.expected.String())
		}
	}
}
