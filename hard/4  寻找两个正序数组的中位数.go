package hard

/*给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2*/

//暴力合并 然后取中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m == 0 && n != 0 {
		t1 := nums2[(n-1)/2]
		t2 := nums2[n/2]
		return float64(t1+t2) / 2
	} else if m != 0 && n == 0 {
		t1 := nums1[(m-1)/2]
		t2 := nums1[m/2]
		return float64(t1+t2) / 2
	} else if m == 0 && n == 0 {
		return 0
	}

	arr := make([]int, 0)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			arr = append(arr, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			arr = append(arr, nums2[j])
			j++
		} else {
			arr = append(arr, nums1[i])
			arr = append(arr, nums2[j])
			i++
			j++
		}
	}
	if i < m {
		arr = append(arr, nums1[i:]...)
	}
	if j < n {
		arr = append(arr, nums2[j:]...)
	}

	t1 := arr[(n+m-1)/2]
	t2 := arr[(n+m)/2]

	return float64(t1+t2) / 2

}

// 二分查找法
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	left := (n + m + 1) / 2
	right := (n + m + 2) / 2
	return float64(getK(nums1, 0, m-1, nums2, 0, n-1, left)+getK(nums1, 0, m-1, nums2, 0, n-1, right)) * 0.5
}

func getK(nums1 []int, start1 int, end1 int, nums2 []int, start2 int, end2 int, k int) int {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1
	if len1 > len2 {
		return getK(nums2, start2, end2, nums1, start1, end1, k)
	}
	if len1 == 0 {
		return nums2[start2+k-1]
	}
	if k == 1 {
		return getMin(nums2[start2], nums1[start1])
	}
	i := start1 + getMin(len1, k/2) - 1
	j := start2 + getMin(len2, k/2) - 1
	if nums1[i] > nums2[j] {
		return getK(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	} else {
		return getK(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
	}
}

func getMin(n1, n2 int) int {
	if n1 > n2 {
		return n2
	} else {
		return n1
	}
}
