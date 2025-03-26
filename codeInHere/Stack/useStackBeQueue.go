package Stack

//232. 用栈实现队列
//请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）
//实现 MyQueue 类：
//void push(x) 将元素 x 推到队列的末尾
//int pop() 从队列的开头移除并返回元素
//int peek() 返回队列开头的元素
//boolean empty() 如果队列为空，返回 true ；否则，返回 false
/*
算法思路:双栈实现队列
1.使用两个栈stack1和stack2：
  - stack1作为输入栈，用于处理push操作
  - stack2作为输出栈，用于处理pop和peek操作
2.push操作：直接压入stack1
3.pop和peek操作：
  - 如果stack2为空，将stack1的所有元素倒序压入stack2
  - 如果stack2不为空，直接从stack2操作
4.empty操作：当两个栈都为空时，队列为空
*/

type MyQueue struct {
	stack1 []int // 输入栈
	stack2 []int // 输出栈
}

// Constructor1 构造函数，初始化两个栈
func Constructor1() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

// Push 将元素x推到队列末尾
func (this *MyQueue) Push(x int) {
	// 直接压入输入栈
	this.stack1 = append(this.stack1, x)
}

// Pop 从队列开头移除并返回元素
func (this *MyQueue) Pop() int {
	// 如果输出栈为空，将输入栈的所有元素倒序压入输出栈
	if len(this.stack2) == 0 {
		this.move()
	}
	// 从输出栈弹出元素
	x := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return x
}

// Peek 返回队列开头的元素
func (this *MyQueue) Peek() int {
	// 如果输出栈为空，将输入栈的所有元素倒序压入输出栈
	if len(this.stack2) == 0 {
		this.move()
	}
	// 返回输出栈顶部元素
	return this.stack2[len(this.stack2)-1]
}

// Empty 如果队列为空，返回true；否则，返回false
func (this *MyQueue) Empty() bool {
	// 当两个栈都为空时，队列为空
	return len(this.stack1) == 0 && len(this.stack2) == 0
}

// move 辅助函数：将stack1的所有元素倒序压入stack2
func (this *MyQueue) move() {
	// 将stack1的元素依次弹出并压入stack2
	for len(this.stack1) > 0 {
		x := this.stack1[len(this.stack1)-1]
		this.stack1 = this.stack1[:len(this.stack1)-1]
		this.stack2 = append(this.stack2, x)
	}
}
