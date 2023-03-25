package main

import "testing"

func TestEx0(t *testing.T) {
	t0 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	t1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	headNode := mergeTwoLists(t0, t1)
	checkNode := headNode
	t.Log("iteratng over the list")
	for {
		if checkNode.Next == nil {
			break
		}
		t.Log("Val:", checkNode.Val)
		checkNode = checkNode.Next
	}

	truVals := []int{1, 1, 2, 3, 4, 4}
	idx := 0
	for {
		if headNode == nil || idx == len(truVals) {
			break
		}
		if headNode.Val != truVals[idx] {
			t.Error("Incorrect node value", truVals[idx], headNode)
			return
		}
		headNode = headNode.Next
		idx++
	}
	if headNode != nil && idx != len(truVals) {
		t.Error("I think there was a mismatch")
		t.Log("HeadNode", headNode)
		t.Log("idx", idx)
		t.Log("True Values", truVals)
	}
}
