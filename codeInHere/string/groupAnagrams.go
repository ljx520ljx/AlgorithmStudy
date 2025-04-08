package string

import (
	"slices"
	"sort"
)

/*字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
*/
/*
算法思路:先排序,再通过hashMap分组
最后把每组的结果给res即可
*/

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, s := range strs {
		t := []byte(s)
		slices.Sort(t)
		sortedS := string(t)
		m[sortedS] = append(m[sortedS], s) // sortedS 相同的字符串分到同一组
	}
	ans := make([][]string, 0, len(m)) // 预分配空间
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func groupAnagrams1(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func groupAnagrams2(strs []string) [][]string {
	var res [][]string
	strMap := make(map[string][]string, len(strs))

	for _, s := range strs {
		// 将字符串转为字符切片并排序，直接使用 string 转回排序后的字符串作为 map 键
		sortedStr := sortString(s)
		strMap[sortedStr] = append(strMap[sortedStr], s)
	}

	// 将结果组装到 res 中
	for _, group := range strMap {
		res = append(res, group)
	}

	return res
}

func sortString(s string) string {
	chars := []byte(s) // 将字符串转为字节切片
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}
