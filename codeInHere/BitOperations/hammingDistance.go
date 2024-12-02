package BitOperations

import "math/bits"

//汉明距离
//两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
//给你两个整数 x 和 y，计算并返回它们之间的汉明距离。
/*
算法思路:异或运算之后,1的数量就是答案
方法一：内置位计数功能,编程语言都内置了计算二进制表达中 1 的数量的函数。
方法二：移位实现位计数:记 s=x⊕y，我们可以不断地检查 s 的最低位，如果最低位为 1，
那么令计数器加一，然后我们令 s 整体右移一位，这样 s 的最低位将被舍去，原本的次低位就变成了新的最低位。
方法三：Brian Kernighan 算法:记 f(x) 表示 x 和 x−1 进行与运算所得的结果（即 f(x)=x&(x−1)），那么 f(x) 恰为 x 删去其二进制表示中最右侧的 1 的结果。
假如x=1001000,x-1=1000111,那么x&x-1=1000000
基于该算法，当我们计算出 s=x⊕y，只需要不断让 s=f(s)，直到 s=0 即可。这样每循环一次，s 都会删去其二进制表示中最右侧的 1，最终循环的次数即为 s 的二进制表示中 1 的数量。
*/

func hammingDistance(x, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func hammingDistance1(x, y int) (ans int) {
	for s := x ^ y; s > 0; s >>= 1 {
		ans += s & 1
	}
	return
}

func hammingDistance3(x, y int) (ans int) {
	for s := x ^ y; s > 0; s &= s - 1 {
		ans++
	}
	return
}
