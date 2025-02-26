package SpecialAlgorithm

// 14. 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
/*
算法思路:选择第一个字符串作为基准，依次和后面的字符串进行纵向比较，
当前前缀下标等于被比较的字符串的长度或者出现了不匹配的字符时，
直接返回当前前缀。
否则到最后这个基准就是最长公共前缀
*/

func longestCommonPrefix(strs []string) string {
	s0 := strs[0]
	for j := range s0 {
		for _, s := range strs {
			if j >= len(s) || s0[j] != s[j] {
				return s0[:j]
			}
		}
	}
	return s0
}
