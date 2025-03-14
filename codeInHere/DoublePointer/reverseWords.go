package DoublePointer

import "strings"

//151. 反转字符串中的单词
/*给你一个字符串 s ，请你反转字符串中 单词 的顺序。
示例 1：输入：s = "the sky is blue"输出："blue is sky the"
并且去除多余的空格*/
/*
算法思路:
先从后往前把每个单词从前往后存在一个字符串切片里,然后遍历切片把每个单词加起来成为答案
在拆分单词的过程中,使用双指针end和index记录单词的开始和结束位置,然后把单词加入切片
*/
func reverseWords(s string) string {
	n := len(s)
	str := make([]string, 0) // 初始化一个空切片
	index := n - 1

	for index >= 0 {
		// 跳过空格
		if s[index] == ' ' {
			index--
			continue
		} else {
			// 记录单词的结束位置
			end := index
			// 找到单词的起始位置
			for index >= 0 && s[index] != ' ' {
				index--
			}
			// 提取单词并加入切片
			str = append(str, s[index+1:end+1])
		}
	}

	// 将切片中的单词拼接成结果字符串
	var res string
	for k, v := range str {
		if k != len(str)-1 {
			res += v + " "
		} else {
			res += v
		}
	}
	return res
	//这上面一整段可以用return strings.Join(str," ")代替
}

/*
使用自带函数的算法思路:
去除首尾空格
将字符串按空格分割成单词数组
反转单词数组
将单词数组拼接成字符串，并用单个空格连接
*/

func reverseWords1(s string) string {
	// 去除首尾空格,制表符 \t、换行符 \n 等可以省略因为Fields足够强大
	s = strings.TrimSpace(s)
	// 将字符串按空格分割成单词数组
	words := strings.Fields(s)
	// 反转单词数组
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// 将单词数组拼接成字符串，并用单个空格连接
	return strings.Join(words, " ")
}
