package linklist

import "fmt"

/*输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
示例 1：
输入：head = [1,3,2]
输出：[2,3,1]
*/
type ListNode struct {
	Next *ListNode
	Val  int
}

func NewNodeList(node *ListNode, val int) *ListNode {
	l := new(ListNode)
	l.Val = val
	l.Next = node
	return l
}

func main() {
	n3 := NewNodeList(nil, 3)
	n2 := NewNodeList(n3, 2)
	n1 := NewNodeList(n2, 1)

	fmt.Println(reversePrint(n1))
}

func reversePrint(head *ListNode) []int {
	arr := make([]int, 0)
	searchLast(head, &arr)
	return arr
}

func searchLast(head *ListNode, arr *[]int) {
	if head != nil {
		searchLast(head.Next, arr)
		*arr = append(*arr, head.Val)
	} else {
		return
	}
}
