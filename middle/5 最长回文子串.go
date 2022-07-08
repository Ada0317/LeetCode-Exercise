package middle

/*给你一个字符串 s，找到 s 中最长的回文子串。*/

func longestPalindrome(s string) string {
	var maxS string
	for i := 0; i < len(s); i++ {
		s1 := getMax(s, i, i)
		s2 := getMax(s, i, i+1)
		if len(s1) > len(maxS) {
			maxS = s1
		}
		if len(s2) > len(maxS) {
			maxS = s2
		}
	}
	return maxS
}

func getMax(s string, i, j int) string {
	l, r := 0, len(s)-1
	for i >= l && j <= r {
		if s[i] == s[j] {
			i--
			j++
		} else {
			break
		}
	}
	return s[i+1 : j]
}
