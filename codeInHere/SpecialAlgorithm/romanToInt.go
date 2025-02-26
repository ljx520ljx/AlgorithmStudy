package SpecialAlgorithm

// RoMan 13. 罗马数字转整数
//罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。把罗马数字转换成整数。
/*
算法思路:
核心点在于处理它的特殊规则,但是这些规则统一起来就是:
x,y是两个相邻的罗马数字
如果x<y,那么就看作-x,否则就是+x
最后把所有的数值加起来就是答案
*/

var RoMan = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	res := 0
	for i := 1; i < len(s); i++ {
		x, y := RoMan[s[i-1]], RoMan[s[i]]
		if x < y {
			res += -x
		} else {
			res += x
		}
	}
	return res + RoMan[s[len(s)-1]]
}
