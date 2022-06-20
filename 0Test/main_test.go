package _Test

import (
	"fmt"
	"leetcodeExe/offer2/tree"
	"testing"
)

func TestName01(t *testing.T) {
	left1 := tree.TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	right1 := tree.TreeNode{
		Val:   2,
		Left:  &left1,
		Right: nil,
	}

	root := tree.TreeNode{
		Val:   1,
		Left:  nil,
		Right: &right1,
	}

	arr := tree.PreOrder_UseStack(&root)
	fmt.Println(arr)
	arr = tree.PreorderTraversal(&root)
	fmt.Println(arr)
}
