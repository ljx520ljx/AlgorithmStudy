package DoublePointer

//165 比较版本号
/*
给你两个 版本号字符串 version1 和 version2 ，请你比较它们。版本号由被点 '.' 分开的修订号组成。修订号的值 是它 转换为整数 并忽略前导零。
比较版本号时，请按 从左到右的顺序 依次比较它们的修订号。如果其中一个版本字符串的修订号较少，则将缺失的修订号视为 0。
*/
/*
算法思路:双指针
用两个指针分别指向两个字符串的起始位置，然后从左到右遍历两个字符串，每次遍历到一个点号时，就将两个指针之间的字符串转换为整数，然后比较两个整数的大小。
在对比两个数的大小进行返回值
注意点:当遇到'.'时，需要将字符串转换为整数，然后index指针要后移一位跳过点号指向下一个数字
*/
func compareVersion(version1 string, version2 string) int {
	m, n := len(version1), len(version2)
	i, j := 0, 0
	for i < m || j < n {
		x := 0
		for ; i < m && version1[i] != '.'; i++ {
			x = x*10 + int(version1[i]-'0')
		}
		i++ //跳过点号
		y := 0
		for ; j < n && version2[j] != '.'; j++ {
			y = y*10 + int(version2[j]-'0')
		}
		j++
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}
