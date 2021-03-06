package tree

/*输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。

假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
*/

// buildTree 递归法
/*方法一：递归
思路

对于任意一颗树而言，前序遍历的形式总是

[ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
即根节点总是前序遍历中的第一个节点。而中序遍历的形式总是

[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
只要我们在中序遍历中定位到根节点，那么我们就可以分别知道左子树和右子树中的节点数目。
由于同一颗子树的前序遍历和中序遍历的长度显然是相同的，因此我们就可以对应到前序遍历的结果中，对上述形式中的所有左右括号进行定位。

这样以来，我们就知道了左子树的前序遍历和中序遍历结果，以及右子树的前序遍历和中序遍历结果，
我们就可以递归地对构造出左子树和右子树，再将这两颗子树接到根节点的左右位置。

细节：
在中序遍历中对根节点进行定位时，一种简单的方法是直接扫描整个中序遍历的结果并找出根节点，
但这样做的时间复杂度较高。我们可以考虑使用哈希表来帮助我们快速地定位根节点。
对于哈希映射中的每个键值对，键表示一个元素（节点的值），值表示其在中序遍历中的出现位置。
在构造二叉树的过程之前，我们可以对中序遍历的列表进行一遍扫描，就可以构造出这个哈希映射。
在此后构造二叉树的过程中，我们就只需要 O(1)O(1) 的时间对根节点进行定位了。
*/

/*前序遍历preorder（根->左->右）的第一个元素是根，设为A
在中序遍历inorder（左->根->右）中找到A，设索引为i，则左子树的长度即为i，那么：
inorder中左子树为：inorder[0:i]（golang的slice操作是左闭右开区间，i指向树根，故不包含i），右子树为：inorder[i+1:]
preorder中左子树为：preorder[1:i+1]（左子树部分长度为i），右子树为：preorder[i+1:]
同理可得递归构建子树*/

func buildTree(preorder []int, inorder []int) *TreeNode { //slice是左闭右开
	inorderMap := make(map[int]int) //[key] index
	for k, v := range inorder {
		inorderMap[v] = k
	}

	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]

	i, ok := inorderMap[rootVal]
	if ok != true {
		return nil
	}

	root := &TreeNode{
		Val: rootVal,
	}
	//重复建立map 导致性能较差
	root.Left = buildTree(preorder[1:i+1], inorder[0:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

func buildTreeBYArr(preorder []int, inorder []int) *TreeNode {
	// terminator, 所有节点构建完成
	if len(preorder) == 0 {
		return nil
	}
	// i:树根的索引，rootval:树根的值
	i, rootval := 0, preorder[0]
	for ; i < len(inorder); i++ {
		if inorder[i] == rootval {
			break
		}
	}
	// 构建新树，左子树部分为：inorder[0:i], preorder[1:i+1]，右子树部分为：inorder[i+1:], preorder[i+1:]
	root := &TreeNode{Val: rootval}
	root.Left = buildTree(preorder[1:i+1], inorder[0:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
