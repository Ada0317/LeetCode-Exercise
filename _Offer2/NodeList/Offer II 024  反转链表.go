package linklist

/*给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
*/

func reverseList(head *ListNode) *ListNode {

	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.Next

		curr.Next = pre

		pre = curr

		curr = next
	}

	return pre
}
