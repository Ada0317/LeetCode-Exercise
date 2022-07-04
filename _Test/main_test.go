package _Test

import (
	"fmt"
	"testing"
)

func Test_02(t *testing.T) {

	//fmt.Println(arr[1:])

}
func TestName01(t *testing.T) {
	//left1 := tree.TreeNode{
	//	Val:   3,
	//	Left:  nil,
	//	Right: nil,
	//}
	//right1 := tree.TreeNode{
	//	Val:   2,
	//	Left:  &left1,
	//	Right: nil,
	//}
	//
	//root := tree.TreeNode{
	//	Val:   1,
	//	Left:  nil,
	//	Right: &right1,
	//}
	//
	//arr := tree.PreOrder_UseStack(&root)
	//fmt.Println(arr)
	//arr = tree.PreorderTraversal(&root)
	//fmt.Println(arr)
	n7 := &ListNode{
		Val:  7,
		Next: nil,
	}
	n3 := &ListNode{
		Val:  3,
		Next: n7,
	}
	n9 := &ListNode{
		Val:  9,
		Next: n3,
	}

	n31 := &ListNode{
		Val:  3,
		Next: nil,
	}
	n6 := &ListNode{
		Val:  6,
		Next: n31,
	}

	l1, l2 := addInList(n9, n6)
	fmt.Println(l1, l2)

}
func addInList(head1 *ListNode, head2 *ListNode) ([]*ListNode, []*ListNode) {
	// write code here
	stack1 := make([]*ListNode, 0)
	stack2 := make([]*ListNode, 0)
	for head1 != nil {
		temp := []*ListNode{head1.Next}
		stack1 = append(temp, stack1...)
		head1 = head1.Next
	}
	for head2 != nil {
		temp := []*ListNode{head2.Next}
		stack1 = append(temp, stack2...)
		head2 = head2.Next
	}
	return stack1, stack2
}

type ListNode struct {
	Val  int
	Next *ListNode
}
