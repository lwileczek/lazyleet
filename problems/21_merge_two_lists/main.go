package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

//ListNode  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		if list2 == nil {
			return list1
		}

		if list1 == nil {
			return list2
		}

		return &ListNode{}
	}

	var root *ListNode
	if list1.Val > list2.Val {
		root = &ListNode{
			Val:  list2.Val,
			Next: mergeTwoLists(list1, list2.Next),
		}
	} else {
		root = &ListNode{
			Val:  list1.Val,
			Next: mergeTwoLists(list1.Next, list2),
		}
	}

	return root
}

//Leetcode said this ran quicker. No recurssion
func bestMergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = ListNode{}

	var p = &dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 == nil {
		p.Next = l2
	}

	if l2 == nil {
		p.Next = l1
	}

	return dummy.Next
}
