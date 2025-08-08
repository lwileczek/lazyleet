package main

import (
	"testing"
)

func TestToLinkedList(t *testing.T) {
	linkedList := FromSlice([]int{2, 4, 3})
	i := linkedList.Int()
	if i != 342 {
		t.Errorf("failed to construct linked list\n\tExpected %d, got %d\n\t%+v", 342, i, linkedList)
	}
}

func TestFromLinkedList(t *testing.T) {
	type TestCase struct {
		List     ListNode
		Expected int
		Msg      string
	}

	testCases := []TestCase{
		{
			List: ListNode{
				Val: 2,
			},
			Expected: 2,
			Msg:      "Single number, easy mode",
		},
		{
			List: ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			Expected: 342,
			Msg:      "First example from problem, [2,4,3] => 342",
		},
	}

	for _, tc := range testCases {
		i := tc.List.Int()
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
	}

	for _, tc := range testCases {
		ans := addTwoNumbers(tc.l1, tc.l2)
		if ans.String() != tc.expected.String() {
			t.Errorf("Failed to add two numbers\n\tgot %s but expected %s\n", ans.String(), tc.expected.String())
		}
	}
}
