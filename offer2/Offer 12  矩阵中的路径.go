package offer2

/*给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。
["A","B","C","E"]
["S","F","C","S"]
["A","D","E","E"]

示例 1：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

示例 2：
输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false
*/

type pair struct {
	X int
	Y int
}

func exist(board [][]byte, word string) bool {

	var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} //左右上下

	rows := len(board)
	cols := len(board[0])
	vis := make([][]bool, rows) //辅助数组标记访问路径
	for i := range vis {
		vis[i] = make([]bool, cols)
	}

	var check func(i int, j int, k int) bool

	check = func(i int, j int, k int) bool { // 行列 和 word索引
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		vis[i][j] = true
		defer func() {
			vis[i][j] = false
		}()
		for _, dir := range directions {
			if newI, newJ := i+dir.X, j+dir.Y; 0 <= newI && newI < rows && 0 <= newJ && newJ <= cols && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i, rowBody := range board {
		for j := range rowBody {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
