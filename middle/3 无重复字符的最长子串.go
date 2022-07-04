package middle

import "strings"

//给定一个字符串 s 请你找出其中不含有重复字符的 最长子串 的长度。

//1 记录左边界 遍历右边界

func lengthOfLongestSubstring(s string) int {
	res := 0
	left, right := -1, 0 //记录左边界  遍历右边界
	for ; right < len(s); right++ {
		if left == -1 {
			left = right
		} else {
			for j := right - 1; j >= left; j-- {
				if s[right] == s[j] {
					left = j + 1
					break
				}
			}
		}
		if right-left+1 > res {
			res = right - left + 1
		}

	}
	return res
}

//只增大的滑动窗口
func lengthOfLongestSubstring2(s string) int {
	start, end := 0, 0
	for i := range s {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index
		}
	}
	return end - start
}
