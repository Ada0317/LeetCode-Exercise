package middle

import "strings"

/*给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
假设二叉树中至少有一个节点。
输入: root = [2,1,3]
输出: 1
*/

//方法一 深度优先搜索 递归法

func findBottomLeftValue1(root *TreeNode) int {
	curHight := 0
	res := root.Val

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, high int) {
		if node == nil {
			return
		}
		high++

		if node.Left != nil {
			dfs(node.Left, high)
		}
		if node.Right != nil {
			dfs(node.Right, high)
		}

		if high > curHight {
			res = node.Val
			curHight = high
		}
	}
	dfs(root, 0)

	return res
}

// 方法2 广度优先搜索
func findBottomLeftValue2(root *TreeNode) int {
	que := make([]*TreeNode, 0)
	que = append(que, root)
	var res int
	for len(que) != 0 {
		node := que[0]
		res = node.Val
		if node.Right != nil {
			que = append(que, node.Right)
		}
		if node.Left != nil {
			que = append(que, node.Left)
		}

		que = que[1:]
	}
	strings.HasPrefix()
	return res
}
