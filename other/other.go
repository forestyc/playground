package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	left := root
	right := root
	// left
	for {
		if left.Left == nil {
			break
		}
		if (left.Left != nil && left.Val <= left.Left.Val) || (right.Right != nil && right.Val >= right.Right.Val) {
			return false
		}
		left = left.Left
	}
	// right
	for {
		if right.Right == nil {
			break
		}
		if (left.Left != nil && left.Val <= left.Left.Val) || (right.Right != nil && right.Val >= right.Right.Val) {
			return false
		}
		right = right.Right
	}
	return true
}

func main() {
	tree := TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}},
	}
	fmt.Println(isValidBST(&tree))
}
