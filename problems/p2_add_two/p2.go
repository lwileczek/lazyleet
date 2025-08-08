package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) elements() []int {
	if n == nil {
		return nil
	}
	s := []int{n.Val}
	s = append(s, n.Next.elements()...)
	return s
}

func (n *ListNode) String() string {
	if n == nil {
		return ""
	}

	s := fmt.Sprintf("%+v", n.elements())

	return s
}

func (n *ListNode) Int() int {
	if n == nil {
		return 0
	}

	v := n.Val + 10*n.Next.Int()
	return v
}

func FromSlice(v []int) *ListNode {
	if len(v) == 0 {
		return nil
	}

	slices.Reverse(v)
	var root *ListNode
	var prev *ListNode
	for _, i := range v {
		root = &ListNode{
			Val:  i,
			Next: prev,
		}
		prev = root
	}

	return root
}

func ToLinkedList(v int) (*ListNode, error) {
	str := strconv.Itoa(v)
	digits := strings.Split(str, "")
	ints := make([]int, len(digits))
	for j, num := range digits {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}

		ints[j] = n
	}

	slices.Reverse(ints)
	node := FromSlice(ints)

	return node, nil
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	t := l1.Int() + l2.Int()
	node, err := ToLinkedList(t)
	if err != nil {
		panic(err)
	}

	return node
}

func main() {
	l1 := &ListNode{Val: 2}
	l2 := &ListNode{Val: 3}
	ans := addTwoNumbers(l1, l2)
	fmt.Println(ans.String())

	linkedList := FromSlice([]int{2, 4, 3})
	fmt.Println(linkedList)
}
