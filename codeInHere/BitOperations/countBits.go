package BitOperations

import "math/bits"

// 338. 比特位计数
/*
算法思路:Brian Kernighan 算法或者直接调用库函数
这题就简单的把从0到n的每个数都进行对1的数量进行计算
*/

func countBits(n int) []int {
	var res []int
	for i := 0; i <= n; i++ {
		count := 0
		num := i
		//num &= num - 1是把num的二进制表示的最低位的1变为0
		for ; num > 0; num &= num - 1 {
			count++
		}
		res = append(res, count)
	}
	return res
}

func countBits1(n int) []int {
	var res []int
	for i := 0; i <= n; i++ {
		res = append(res, bits.OnesCount(uint(i)))
	}
	return res
}

func countBits2(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		//i>>1是把i右移一位,i%2是取i的二进制表示的最低位为0或1
		bits[i] = bits[i>>1] + i%2
	}
	return bits
}

func countBits3(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		//i&(i-1)是把i的二进制表示的最低位的1变为0,再把删掉的1加回去
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}

func countBits4(n int) []int {
	bits := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		//i-highBit是把i的二进制表示的最高位的1变为0,再把删掉的1加回去
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}
