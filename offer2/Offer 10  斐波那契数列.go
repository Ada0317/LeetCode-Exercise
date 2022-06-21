package offer2

//剑指 Offer 10- I. 斐波那契数列
//写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
//
//F(0) = 0, F(1)= 1
//F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
//斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
//
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

/*
注意取模和初始条件即可

*/
func fib(n int) int {
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	a := 0
	b := 1

	for i := 2; i <= n; i++ {
		res := (a + b) % mod
		a = b
		b = res
	}
	return b

}

//剑指 Offer 10- II. 青蛙跳台阶问题
//一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

func numWays(n int) int {
	const mod = 1e9 + 7
	if n <= 1 {
		return 1
	}

	num1, num2 := 1, 1
	for i := 2; i <= n; i++ {
		num2, num1 = (num1+num2)%mod, num2

	}
	return num2
}
