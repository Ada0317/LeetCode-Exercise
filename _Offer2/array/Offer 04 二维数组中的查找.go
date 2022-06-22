package array

import (
	"fmt"
)

/*在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

示例:

现有矩阵 matrix (矩阵) 如下：

[
[1,   4,  7, 11, 15],
[2,   5,  8, 12, 19],
[3,   6,  9, 16, 22],
[10, 13, 14, 17, 24],
[18, 21, 23, 26, 30]
]
*/

/*
方法一:
暴力
如果不考虑二维数组排好序的特点，则直接遍历整个二维数组的每一个元素，判断目标值是否在二维数组中存在。

依次遍历二维数组的每一行和每一列。如果找到一个元素等于目标值，则返回 true。如果遍历完毕仍未找到等于目标值的元素，则返回 false。*/

/*时间复杂度：O(nm)。二维数组中的每个元素都被遍历，因此时间复杂度为二维数组的大小。
空间复杂度：O(1)。*/

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
	}

	exist := method01(matrix, 5)
	fmt.Println(exist)
	exist = method02(matrix, 5)
	fmt.Println(exist)
}

func method01(matrix [][]int, target int) bool {

	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if target == matrix[i][j] {
				return true
			}
		}
	}

	return false

}

/*
方法二：线性查找
由于给定的二维数组具备每行从左到右递增以及每列从上到下递增的特点，当访问到一个元素时，可以排除数组中的部分元素。

从二维数组的右上角开始查找。如果当前元素等于目标值，则返回 true。如果当前元素大于目标值，则移到左边一列。如果当前元素小于目标值，则移到下边一行。

可以证明这种方法不会错过目标值。如果当前元素大于目标值，说明当前元素的下边的所有元素都一定大于目标值，因此往下查找不可能找到目标值，往左查找可能找到目标值。如果当前元素小于目标值，说明当前元素的左边的所有元素都一定小于目标值，因此往左查找不可能找到目标值，往下查找可能找到目标值。

*/

/*时间复杂度：O(n+m)。访问到的下标的行最多增加 n 次，列最多减少 m 次，因此循环体最多执行 n + m 次。
空间复杂度：O(1)。*/

func method02(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows := len(matrix)
	column := len(matrix[0])
	i := 0
	j := column - 1
	for i < rows && j >= 0 {
		num := matrix[i][j]
		if target < num {
			j--
		} else if target > num {
			i++
		} else {
			return true
		}
	}
	return false
}
