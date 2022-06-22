package middle

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail *ListNode
	var head *ListNode

	carry := 0

	for l1 != nil || l2 != nil {
		num1, num2 := 0, 0
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}
		temp := num1 + num2 + carry
		if temp > 9 {
			temp = temp - 10
			carry = 1
		} else {
			carry = 0
		}
		if head == nil {
			head = &ListNode{
				Val: temp,
			}
			tail = head
		} else {
			tail.Next = &ListNode{Val: temp}
			tail = tail.Next
		}

		/*if head == nil {
			head = &ListNode{
				Val: temp,
			}
			tail = head.Next
		} else {
			tail = &ListNode{Val: temp}
			tail = tail.Next
		}*/ //这样是错误的 因为链表断开了 只会返回第一个节点

	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return head
}

//优化的方法 减少创建列表
func addTwoNumbersNew(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := l1
	var pre *ListNode

	for l1 != nil && l2 != nil {
		temp := carry + l1.Val + l2.Val
		if temp >= 10 {
			l1.Val = temp - 10
			carry = 1
		} else {
			l1.Val = temp
			carry = 0
		}

		pre = l1
		l1 = l1.Next
		l2 = l2.Next
	}

	if l2 != nil {
		pre.Next = l2 //重新调整l1的指向
		l1 = pre.Next
	}

	for l1 != nil {
		temp := carry + l1.Val
		if temp >= 10 {
			l1.Val = temp - 10
			carry = 1
		} else {
			l1.Val = temp
			carry = 0
		}
		pre = l1
		l1 = l1.Next
	}

	if carry > 0 {
		pre.Next = &ListNode{Val: 1}
	}
	return head
}
