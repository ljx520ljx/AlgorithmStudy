package DoublePointer

//最长回文字符串
/*
算法思路:中心扩展算法+双指针
中心扩展：
回文的中心可能是一个字符，也可能是两个相邻的字符。
我们对每个字符和每对相邻字符尝试作为中心，向两边扩展，找到最大的回文子串。
双指针扩展：
设定两个指针 left 和 right，从中心出发，向左和向右扩展，直到找到最大的回文区间。
每次扩展之后，我们会记录最长的回文子串的长度和起始位置。
时间复杂度：
该算法的时间复杂度是 O(n^2)，因为我们需要对每个字符和字符对作为中心进行一次线性扫描。
由于没有使用额外的复杂数据结构，空间复杂度是 O(1)。
*/

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	// 定义一个辅助函数，用于从中心向外扩展
	expandAroundCenter := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		// 以单个字符为中心
		expandAroundCenter(i, i)
		// 以两个相邻字符为中心
		expandAroundCenter(i, i+1)
	}

	return s[start : start+maxLen]
}

//自己写的史,超长复杂度
//func longestPalindrome(s string) string {
//	var maxs string
//	charmap:=make(map[byte][]int)
//	s1:=[]byte(s)
//	for i,char:=range s1{
//		charmap[char]=append(charmap[char],i)
//	}
//	for j:=range s1{
//		a:=charmap[s1[j]]
//		for _,i:=range a{
//			right:=i
//			left:=j
//			length:=right-left+1
//			for right>left&&s1[right]==s1[left]{
//				right--
//				left++
//			}
//			if left>=right&&length>len(maxs){
//				maxs=s[j:i+1]
//			}
//		}
//	}
//	return maxs
//}
