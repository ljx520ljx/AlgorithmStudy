package SlidingWindow

// 找到给定字符串中无重复字符的最长子串长度

/* 算法思路:
滑动窗口：我们使用滑动窗口的技术，用 left 和 right 两个指针来定义窗口的左右边界,
遍历整个字符串。窗口始终保持无重复字符的状态。
哈希表：charMap 是一个哈希表（即 map），用于存储字符和它们最近出现的位置。
这样可以快速找到某个字符上次出现的索引。
移动窗口：对于每个新的字符（即 s[right]），
检查该字符是否在当前窗口（即从 left 到 right）中已经存在。
如果存在，说明有重复字符，此时将 left 指针移动到重复字符的下一个位置。
更新字符在 charMap 中的最新位置为 right。
更新最大长度：在每次迭代中，计算当前窗口的长度 right - left + 1，
并将其与 maxLength 比较，保留较大值。
*/

// 数据结构:采用map[byte]int.

func lengthOfLongestSubstring(s string) int {
	// 使用 map 存储字符的最后一次出现的位置
	charMap := make(map[byte]int)
	maxLength := 0 // 保存无重复字符子串的最大长度
	left := 0      // 滑动窗口的左边界

	// 遍历字符串的每个字符，right 是滑动窗口的右边界
	for right := 0; right < len(s); right++ {
		// 检查当前字符 s[right] 是否在当前窗口中出现过
		if index, ok := charMap[s[right]]; ok && index >= left { // 如果找到重复字符，并且它的位置在当前窗口内，
			// 则更新 left 指针，跳过重复字符
			left = index + 1
		}
		// 将当前字符及其位置存入 map 中，更新该字符的最后出现位置
		charMap[s[right]] = right

		// 更新最大长度，当前窗口长度为 right - left + 1
		if maxLength < right-left+1 {
			maxLength = right - left + 1
		}
	}

	// 返回最大长度
	return maxLength
}

//理解题后自己重写的

//func lengthOfLongestSubstring(s string) int {
//	charMap := make(map[rune]int)
//	maxlength, left := 0, 0
//	for right, char := range s {
//		if index, ok := charMap[char]; ok && index >= left {
//			left = index + 1
//		}
//		charMap[char] = right
//		if maxlength < right-left+1 {
//			maxlength = right - left + 1
//		}
//	}
//	return maxlength
//}
