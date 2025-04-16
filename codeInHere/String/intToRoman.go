package String

//12. 整数转罗马数字
/*
算法思路:打表法
把个十百千位的罗马数字都列出来，然后用一个二维数组把它们存起来，
最后把每一位的数字对应的罗马数字拼起来即可。
*/

var R = [4][10]string{
	{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}, // 个位
	{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}, // 十位
	{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}, // 百位
	{"", "M", "MM", "MMM"}, // 千位
}

func intToRoman(num int) string {
	return R[3][num/1000%10] + R[2][num/100%10] + R[1][num/10%10] + R[0][num%10]
}
