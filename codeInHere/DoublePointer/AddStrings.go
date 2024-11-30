package DoublePointer

//字符串相加:给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
/*
算法思路:
初始化指针和变量：
用两个指针分别指向 num1 和 num2 的末尾（表示最低位）。
初始化一个变量 carry 表示进位，初始值为 0。
用一个字符串数组 result 保存计算结果（逆序存储每一位计算结果，最后再反转）。
逐位相加：
从低位到高位，取出两个字符串对应位置的数字，如果某个字符串的指针越界，则该位置的数字视为 0。
将两个数字与当前进位 carry 相加，计算出当前位的结果 digit 和新的进位 carry。
把 digit 转成字符后加入到 result 中。
处理剩余的进位：
如果最后还有进位，需要把它加入 result。
结果反转：
将 result 数组反转后拼接成字符串，得到最终结果。
*/

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	var res []byte
	for i >= 0 || j >= 0 || carry > 0 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
			j--
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		digit := sum % 10
		//这里byte(digit)+'0'是ASCII码的值相加
		res = append(res, byte(digit)+'0')
	}
	//反转res
	for left, right := 0, len(res)-1; left < right; left, right = left+1, right-1 {
		res[left], res[right] = res[right], res[left]
	}
	return string(res)
}
