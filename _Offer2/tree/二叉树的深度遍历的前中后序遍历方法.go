package tree

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrder(node *TreeNode, arr *[]int) {
	if node != nil {
		*arr = append(*arr, node.Val)
		PreOrder(node.Left, arr)

		PreOrder(node.Right, arr)
	} else {
		return
	}
}

func MiddleOrder(node *TreeNode, arr *[]int) {
	if node != nil {
		MiddleOrder(node.Left, arr)
		*arr = append(*arr, node.Val)
		MiddleOrder(node.Right, arr)
	} else {
		return
	}
}

func BackOrder(node *TreeNode, arr *[]int) {
	if node != nil {
		BackOrder(node.Left, arr)
		BackOrder(node.Right, arr)
		*arr = append(*arr, node.Val)
	} else {
		return
	}
}

// PreOrder_UseStack 非递归方法  用到额外数据结构栈  这里直接使用container.list双向链表进行栈模拟
func PreOrderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	pnode := root
	l := list.New() //初始化双向链表
	for pnode != nil || l.Len() != 0 {
		if pnode != nil {
			arr = append(arr, pnode.Val)
			l.PushBack(pnode)
			pnode = pnode.Left
		} else {
			back := l.Back()
			l.Remove(back)
			node, ok := back.Value.(*TreeNode)
			if ok != true {
				panic("转换出错")
			}
			pnode = node.Right
		}
	}
	return arr
}

func MidOrderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	l := list.New()
	pnode := root
	for pnode != nil || l.Len() != 0 {
		if pnode != nil {
			l.PushBack(pnode)
			pnode = pnode.Left
		} else {
			back := l.Back()
			l.Remove(back)
			node, ok := back.Value.(*TreeNode)
			if ok != true {
				panic("转换出错")
			}
			arr = append(arr, node.Val)
			pnode = node.Right
		}
	}
	return arr
}

// 广度优先遍历 利用队列先进先出的特点 先把根节点放入队列 然后把左右子节点放进队列 循环判断队列是否为空就好

func LevelTraversal(root *TreeNode) {
	arr := make([]int, 0)
	l := list.New()
	l.PushBack(root)
	for l.Len() != 0 {
		//出队一个
		front := l.Front()
		node := front.Value.(*TreeNode)
		arr = append(arr, node.Val)

		if node.Left != nil {
			l.PushBack(node.Left)
		}

		if node.Right != nil {
			l.PushBack(node.Right)
		}
	}
}

// 深度优先遍历  前序

func DepthTraversal(root *TreeNode) {
	arr := make([]int, 0)
	l := list.New()
	l.PushBack(root)
	for l.Len() != 0 {
		//出栈
		front := l.Back()
		l.Remove(front)
		node := front.Value.(*TreeNode)

		arr = append(arr, node.Val)

		if node.Right != nil {
			l.PushBack(node.Right)
		}

		if node.Left != nil {
			l.PushBack(node.Left)
		}
	}
}
