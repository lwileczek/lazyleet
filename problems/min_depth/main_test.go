package main

import (
	"testing"
)

func TestTree(t *testing.T) {
	t.Parallel()

	ans := 2
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
	if count != ans {
		t.Fail()
	}
}

func TestSecondTree(t *testing.T) {
	t.Parallel()
	ans := 5
	tree := TreeNode{
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

	count := minDepth(&tree)
	if count != ans {
		t.Fail()
	}
}

func TestThirdTree(t *testing.T) {
	t.Parallel()
	ans := 3
	tree := TreeNode{
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
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  0,
				Left: &TreeNode{},
				Right: &TreeNode{
					Val: 30,
					Left: &TreeNode{
						Val:   29,
						Left:  &TreeNode{},
						Right: &TreeNode{},
					},
					Right: &TreeNode{
						Val:   31,
						Left:  &TreeNode{},
						Right: &TreeNode{},
					},
				},
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}

	count := minDepth(&tree)
	if count != ans {
		t.Fail()
	}
}
