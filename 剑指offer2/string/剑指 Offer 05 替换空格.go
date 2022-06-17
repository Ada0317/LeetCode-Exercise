package main

import "fmt"

/*请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."
限制：
0 <= s 的长度 <= 10000*/

func main() {
	space := replaceSpace("法国 的风格 happy.")
	fmt.Println(space)
}

func replaceSpace(s string) string { //append 函数 追加切片的时候 要加上 ...
	runes := []rune(s)
	buf := make([]rune, 0) //会初始化充满零值 所以这里设置容量为0
	for _, v := range runes {
		if v == ' ' {
			buf = append(buf, []rune("%20")...)
			continue
		}
		buf = append(buf, v)
	}
	s = string(buf)
	return s
}
