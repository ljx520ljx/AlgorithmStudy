package Backtrace

//电话号码的字母组合
//给定一个仅包含数字2-9的字符串,返回所有它能表示的字母组合答案可以按任意顺序返回。
//给出数字到字母的映射与电话按键相同注意,1不对应任何字母。
/*
算法思路:回溯算法
映射表：创建一个映射表 phoneMap，将数字 2-9 映射到对应的字母组合,例如'2':"abc",'3':"def"。
处理边界条件：如果输入字符串为空（即没有数字），直接返回一个空数组。
回溯算法：
使用回溯法动态生成所有可能的字母组合。
递归的每一步对应于处理输入字符串中的一个数字，获取该数字对应的字母列表。
遍历这些字母，并将当前字母添加到组合中，递归处理下一个数字。
如果组合的长度与输入字符串长度相等，表示生成了一个完整组合，将其加入结果数组。
结束条件：当递归深度等于输入字符串长度时，终止递归并保存当前组合。
结果输出：最终返回结果数组 result。
*/

var phoneMap map[byte]string = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

var combinations []string

// GPT
func letterCombinations(digits string) []string {
	// 定义电话号码数字到字母的映射
	//phoneMap := map[byte]string{
	//	'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno",
	//	'7': "pqrs", '8': "tuv", '9': "wxyz",
	//}
	// 如果输入为空，返回空数组
	if len(digits) == 0 {
		return []string{}
	}
	// 结果数组，用于存储最终的所有组合
	var result []string

	// 辅助函数，用于递归生成组合
	var backtrack func(index int, combination string)
	backtrack = func(index int, combination string) {
		// 如果组合长度等于输入数字的长度，存入结果
		if index == len(digits) {
			result = append(result, combination)
			return
		}
		// 获取当前数字对应的字母
		letters := phoneMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			// 递归生成下一个字母组合
			backtrack(index+1, combination+string(letters[i]))
		}
	}
	// 从索引0开始递归
	backtrack(0, "")
	return result
}

// 官解
func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		letters := phoneMap[digits[index]]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}
