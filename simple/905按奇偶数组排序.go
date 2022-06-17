package simple

//双指针头尾赋值法  双指针指向新建数字的头尾  然后遍历一次数组  O(n)时间复杂度  O(1)空间复杂度
func sortArrayByParity01(nums []int) []int {
	start, end := 0, len(nums)-1
	tempArr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		//if start < end {
		if nums[i]%2 == 0 {
			tempArr[start] = nums[i]
			start++
		} else {
			tempArr[end] = nums[i]
			end--
		}
		//}
	}
	return tempArr
}

//双指针头尾交换法  双指针指向原来数组的头尾  然后往中间遍历一次数组  遇到不合的就交换  O(n)时间复杂度  O(1)空间复杂度
func sortArrayByParity02(nums []int) []int {
	start, end := 0, len(nums)-1
	for start <= end {
		for start < end && nums[start]%2 == 0 {
			start++
		}
		for start < end && nums[end]%2 != 0 {
			end--
		}
		if start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}

	}
	return nums
}
