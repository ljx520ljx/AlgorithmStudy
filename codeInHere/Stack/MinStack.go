package Stack

/*实现最小栈
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
实现 MinStack 类:
MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。
*/
/*
实现思路:
主栈 是正常的栈，记录每一个被推入栈的元素。
辅助栈 记录栈中最小值的历史。对于每个被推入主栈的元素，
辅助栈会保存当前最小值。如果新插入的值小于或等于当前最小值，
那么就将它也推入辅助栈。这样，辅助栈的栈顶始终保存着当前栈中的最小值。
*/
type MinStack struct {
	stack    []int // 主栈
	minStack []int // 辅助栈
}

// Constructor 初始化栈
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

// Push 将元素 val 推入栈中
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val) // 推入主栈
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val) // 推入辅助栈，保持最小值
	}
}

// Pop 移除栈顶元素
func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	topElem := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1] // 弹出主栈元素

	// 如果弹出的元素是当前的最小元素，也弹出辅助栈的栈顶元素
	if topElem == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

// Top 获取栈顶元素
func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return -1 // 如果栈为空，返回一个非法值
}

// GetMin 获取栈中的最小元素
func (this *MinStack) GetMin() int {
	if len(this.minStack) > 0 {
		return this.minStack[len(this.minStack)-1]
	}
	return -1 // 如果栈为空，返回一个非法值
}
