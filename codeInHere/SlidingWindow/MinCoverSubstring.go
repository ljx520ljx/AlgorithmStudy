package SlidingWindow

// 最小覆盖子串:给你一个字符串 s 、一个字符串 t 。
//返回 s 中涵盖 t 所有字符的最小子串。
//如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
/*
算法思路:滑动窗口算法结合哈希表
需求统计：首先我们要统计目标字符串 t 中每个字符的出现次数，使用一个哈希表 need 来存储该需求。
窗口统计：在源字符串 s 中建立一个滑动窗口，用另一个哈希表 window 统计当前窗口内每个字符的出现次数。
滑动窗口移动：
使用两个指针 left 和 right，其中 right 指向当前窗口的右边界。
每次向右移动 right，将当前字符添加到窗口统计中。
判断窗口中的字符是否满足 need 的要求，
如果满足要求，则尝试收缩窗口，通过移动 left 来缩小窗口的范围，直到窗口不再满足 need 的需求。
更新最小覆盖子串：在满足要求的窗口内，
如果当前窗口的长度比记录的最小长度短，则 更新最小覆盖子串的位置和长度。
终止条件：当 right 遍历完字符串 s 时，算法结束，返回最小覆盖子串。
*/
func minWindow(s string, t string) string {
	// 统计t中的字符需求
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	// 初始化窗口和相关变量
	window := make(map[byte]int)
	left, right := 0, 0
	valid := 0
	start, length := 0, len(s)+1

	for right < len(s) {
		// 将s[right]包含进窗口
		c := s[right]
		right++

		// 更新窗口中的数据
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}
			// 移除s[left]，即窗口左边界右移
			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	// 返回结果
	if length == len(s)+1 {
		return ""
	}
	return s[start : start+length]
}
