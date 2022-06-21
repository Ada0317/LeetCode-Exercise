package string

import (
	"fmt"
)

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

/*时间复杂度：O(n)。遍历字符串 s 一遍。
空间复杂度：O(n)。额外创建字符数组*/

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

/*
tips：使用append往切片中追加另一个切片的时候 要在第二个参数(追加的切片)后加上...  否则语法错误
	rune类型是int32的别名 相当于4个byte
	byte类型是uint8的别名

*/
