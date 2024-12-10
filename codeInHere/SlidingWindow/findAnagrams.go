package SlidingWindow

import "slices"

//找到字符串中所有的字母异位词
//给定两个字符串 s 和 p，找到 s 中所有 p 的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
/*
算法思路:
方法一:把p转换为字节切片后排序p1,后面遍历s,记p的长度为n,索引为i,取s[i:i+n]转换为字节切片后排序s1
最后string(s1)==string(p1)就把索引i添加到结果中,遍历完成即完成.但此方法时间复杂度高,0(n m log m)
方法二:构造滑动窗口
字符串异位词长度与目标字符串相同,在长字符串中构造等长滑动窗口，维护窗口内字母数量。当窗口内字母数量与目标字符串字母数量一致时，该窗口为目标字符串异位词
数组存储长字符串和滑动窗口中每种字母数量。若长字符串长度小于目标字符串长度，单独处理（因无法构造等长窗口）
统计目标字符串字母数量，同时统计长字符串对应位置字母数量。若此时两者字母数量相同，记录当前窗口起始位置
滑动窗口时，先减少窗口左侧字母数量，增加右侧新进入字母数量。若此时窗口内字母数量与目标字符串字母数量相同，记录当前窗口起始位置。
*/

func findAnagrams(s string, p string) []int {
	n := len(p)
	p1 := []byte(p)
	slices.Sort(p1)
	var res []int
	for i := 0; i < len(s); i++ {
		if i+n > len(s) {
			break
		}
		s1 := []byte(s[i : i+n])
		slices.Sort(s1)
		if string(s1) == string(p1) {
			res = append(res, i)
		}
	}
	return res
}

func findAnagrams1(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}

// 引入变量 differ 来记录当前窗口与字符串 p 中数量不同的字母的个数，只需要判断 differ 是否为零即可。
func findAnagrams2(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	count := [26]int{}
	for i, ch := range p {
		count[s[i]-'a']++
		count[ch-'a']--
	}

	differ := 0
	for _, c := range count {
		if c != 0 {
			differ++
		}
	}
	if differ == 0 {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		if count[ch-'a'] == 1 { // 窗口中字母 s[i] 的数量与字符串 p 中的数量从不同变得相同
			differ--
		} else if count[ch-'a'] == 0 { // 窗口中字母 s[i] 的数量与字符串 p 中的数量从相同变得不同
			differ++
		}
		count[ch-'a']--
		if count[s[i+pLen]-'a'] == -1 { // 窗口中字母 s[i+pLen] 的数量与字符串 p 中的数量从不同变得相同
			differ--
		} else if count[s[i+pLen]-'a'] == 0 { // 窗口中字母 s[i+pLen] 的数量与字符串 p 中的数量从相同变得不同
			differ++
		}
		count[s[i+pLen]-'a']++

		if differ == 0 {
			ans = append(ans, i+1)
		}
	}
	return
}
