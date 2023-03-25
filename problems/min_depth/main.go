// https://leetcode.com/problems/minimum-depth-of-binary-tree/
/** Given a binary tree, find its minimum depth.
The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
Note: A leaf is a node with no children.
*/
// submission: https://leetcode.com/submissions/detail/832082105/
package main

import (
	"fmt"
)

//TreeNode Example binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	leaf9 := TreeNode{Val: 9}
	leaf15 := TreeNode{Val: 15}
	leaf7 := TreeNode{Val: 7}
	leaf20 := TreeNode{
		Val:   20,
		Left:  &leaf15,
		Right: &leaf7,
	}
	tree := TreeNode{
		Val:   3,
		Left:  &leaf9,
		Right: &leaf20,
	}
	count := minDepth(&tree)
	fmt.Println(count)
	tree1 := TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
		},
	}

	secondTry := minDepth(&tree1)
	fmt.Println("Second tree parse", secondTry)
}

// Really similar to what I had but they call the recurssion first. How does it run if no Left?
func minDepth(root *TreeNode) (result int) {
	if root == nil {
		return 0
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if root.Left != nil && root.Right == nil {
		return l + 1
	}

	if root.Left == nil && root.Right != nil {
		return r + 1
	}

	result = l
	if r < l {
		result = r
	}
	return result + 1
}

// A method many took that had quick results
func minDepthQueue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// BFS
	queue := []*TreeNode{}
	queue = append(queue, root)

	level := 1
	for len(queue) > 0 {
		sz := len(queue)

		for i := 0; i < sz; i++ {
			// pop left
			node := queue[0]
			queue = queue[1:]

			// find leaf node
			if node.Left == nil && node.Right == nil {
				return level
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
	}

	return 0
}
