package array

import "fmt"

/*找出数组中的重复数字*/

func main() {
	type matrix [][]int
	arr := []int{1, 2, 3, 4, 0, 01}
	number := findRepeatNumber01(arr)
	fmt.Println(number)
	number = findRepeatNumber02(arr)
	fmt.Println(number)
}

/*方法一：
哈希Map
利用数据结构特点，容易想到使用哈希表（Set）记录数组的各个数字，当查找到重复数字则直接返回。
*/

/*时间复杂度 O(N) ： 遍历数组使用 O(N) ，HashSet 添加与查找元素皆为 O(1) 。
空间复杂度 O(N)： HashSet 占用 O(N) 大小的额外空间。
*/
func findRepeatNumber01(nums []int) int {
	m := make(map[int]int)
	for k, v := range nums {
		if _, exist := m[k]; !exist {
			m[k] += 1
		} else {
			return v
		}
	}
	return -1
}

/*方法二：原地交换
题目说明尚未被充分使用，即 在一个长度为 n 的数组 nums 里的所有数字都在 0 ~ n-1 的范围内 。
此说明含义：数组元素的 索引 和 值 是 一对多 的关系。
因此，可遍历数组并通过交换操作，使元素的 索引 与 值 一一对应（即 nums[i] = inums[i]=i ）。
因而，就能通过索引映射对应的值，起到与字典等价的作用。

复杂度分析：
时间复杂度 O(N) ： 遍历数组使用 O(N) ，每轮遍历的判断和交换操作使用 O(1) 。
空间复杂度 O(1) ： 使用常数复杂度的额外空间。

*/

func findRepeatNumber02(nums []int) int {
	for k := 0; k < len(nums); {
		if nums[k] == k {
			k++
			continue
		} else {
			if nums[nums[k]] == nums[k] {
				return nums[k]
			} else {
				nums[nums[k]], nums[k] = nums[k], nums[nums[k]]
			}
		}
	}
	return -1

}
