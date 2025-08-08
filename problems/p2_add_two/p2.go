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

func (n *ListNode) NumString() string {
	if n == nil {
		return ""
	}

	s := ToList(n)
	slices.Reverse(s)
	digits := make([]string, len(s))
	for i, num := range s {
		digits[i] = strconv.Itoa(num)
	}
	return strings.Join(digits, "")
}

func (n *ListNode) String() string {
	if n == nil {
		return ""
	}

	s := fmt.Sprintf("%+v", ToList(n))

	return s
}

func ToList(n *ListNode) []int {
	if n == nil {
		return nil
	}
	s := []int{n.Val}
	s = append(s, ToList(n.Next)...)
	return s
}

func ToInt(n *ListNode) int {
	if n == nil {
		return 0
	}

	v := n.Val + 10*ToInt(n.Next)
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

func addOptionals(l1 *ListNode, l2 *ListNode) int {
	if l1 != nil && l2 != nil {
		return l1.Val + l2.Val
	}

	if l1 == nil {
		return l2.Val
	}

	return l1.Val
}

func addWithCarry(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry == 0 {
			return nil
		}
		return &ListNode{Val: carry}
	}

	v := addOptionals(l1, l2) + carry
	c := 0
	if v > 9 {
		c = 1
		v = v - 10
	}

	var nl *ListNode
	var nr *ListNode
	if l1 != nil {
		nl = l1.Next
	}
	if l2 != nil {
		nr = l2.Next
	}

	sum := &ListNode{
		Val:  v,
		Next: addWithCarry(nl, nr, c),
	}

	return sum
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := addWithCarry(l1, l2, 0)
	return sum
}

func main() {
	l1 := &ListNode{Val: 2}
	l2 := &ListNode{Val: 3}
	ans := addTwoNumbers(l1, l2)
	fmt.Println(ans)

	linkedList := FromSlice([]int{2, 4, 3})
	fmt.Println(linkedList)
}
