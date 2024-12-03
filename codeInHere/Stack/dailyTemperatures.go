package Stack

/*
每日温度
给定一个整数数组 temperatures,表示每天的温度,返回一个数组answer,
其中answer[i]是指对于第 天,下一个更高温度出现在几天后。
如果气温在这之后都不会升高,请在该位置用0来代替。
*/
/*
算法思路：使用栈来存储“待处理”的温度
栈的作用：
我们利用栈来追踪那些“尚未找到下一个更高温度”的天数索引。
栈存放的是温度数组的索引,而不是温度本身。
这使得我们可以在后续的遍历中通过比较当前温度与栈顶温度来确定答案。
遍历过程：
我们遍历整个温度数组。
对于每一天的温度，判断当前温度是否比栈顶的温度高。如果是，说明我们找到了栈顶温度的下一个更高温度。
如果当前温度大于栈顶温度：
弹出栈顶元素，计算当前天数与栈顶天数的差值，这个差值就是栈顶天数的答案（即多少天后温度会更高）。
继续比较下一个栈顶温度。
如果当前温度不大于栈顶温度：
将当前天数的索引压入栈中，表示这个天数的温度还没有找到更高温度的日子。
*/

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	var stack []int
	for i := 0; i < len(temperatures); i++ {

		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			res[idx] = i - idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
