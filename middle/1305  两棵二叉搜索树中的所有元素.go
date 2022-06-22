//给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。
//```
package middle

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode) (res []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}

func getAllElements(root1, root2 *TreeNode) []int {
	nums1 := inorder(root1)
	nums2 := inorder(root2)

	p1, n1 := 0, len(nums1)
	p2, n2 := 0, len(nums2)
	merged := make([]int, 0, n1+n2)
	for {
		if p1 == n1 {
			return append(merged, nums2[p2:]...)
		}
		if p2 == n2 {
			return append(merged, nums1[p1:]...)
		}
		if nums1[p1] < nums2[p2] {
			merged = append(merged, nums1[p1])
			p1++
		} else {
			merged = append(merged, nums2[p2])
			p2++
		}
	}
}

//```
