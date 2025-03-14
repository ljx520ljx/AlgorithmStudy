package DoublePointer

//回文子串
//给你一个字符串 s ，请你统计并返回这个字符串中回文子串的数目。
/*
算法思路:中心拓展+双指针
中心扩展法：对于每个位置 i，我们尝试以这个位置为回文的中心来扩展字符串。回文可以分为两种类型：
以一个字符为中心的回文子串。
以两个字符为中心的回文子串。
对于每个 i，将 l 和 r 指向回文子串的左右边界
在 l >= 0 且 r < n（字符串范围内）且 s[l] == s[r] 的条件下，不断扩展，直到不能再扩展为止。
*/

func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	//i表示有多少个回文中心
	for i := 0; i < 2*n-1; i++ {
		//i为偶数,回文中心是单字符,为奇数,回文中心是双字符
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}
