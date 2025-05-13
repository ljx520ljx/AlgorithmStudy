package Greedy

//678. 有效的括号字符串
//给定一个只包含三种字符的字符串:'(',')'和'*'，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则:
//任何左括号'('必须有相应的右括号')'。
//任何右括号')'必须有相应的左括号'('。
//左括号'('必须在对应的右括号之前')'。
//'*'可以被视为单个右括号')'，或单个左括号'('，或一个空字符串""。
//一个空字符串也被视为有效字符串。
/*
算法思路:贪心算法
使用两个变量 minleft 和 maxleft 来表示当前未匹配的左括号的数量范围。
遍历字符串中的每个字符：
如果遇到左括号 '('，则 minleft 和 maxleft 都加 1。
如果遇到右括号 ')'，则 minleft 和 maxleft 都减 1。
如果遇到星号 '*'，则 maxleft 加 1，minleft 减 1（因为 '*' 可以被视为空字符串或左括号）。
因为 '*' 可以被视为右括号或左括号，而minleft表示最小数量所以--,而maxleft表示最大数量所以++
如果 maxleft 小于 0，则说明右括号的数量超过了左括号和星号的数量，返回 false。
如果 minleft 小于 0，则将其重置为 0，因为左括号的数量不能为负数。
最终，如果 minleft 等于 0，则说明所有左括号都有对应的右括号，返回 true；否则返回 false。
这里当minleft<0时,*可以被视为""
*/

func checkValidString(s string) bool {
	minleft, maxleft := 0, 0
	for _, ch := range s {
		switch ch {
		case '(':
			minleft++
			maxleft++
		case ')':
			maxleft--
			minleft = max(0, minleft-1)
			if maxleft < 0 {
				return false
			}
		case '*':
			maxleft++
			minleft = max(0, minleft-1)
		}
	}
	return minleft == 0
}
