package Queue

import "math"

// 239. 滑动窗口最大值
/*
算法思路:单调队列 (Monotonic Queue)
1. 维护一个双端队列 (deque)，队列中存储的是数组元素的【索引】。
2. 队列中的索引对应的元素值【从队首到队尾单调递减】。
3. 窗口滑动过程：
   a. 入队 (Enqueue): 当新元素进入窗口时，从队尾移除所有小于等于新元素值的索引，然后将新元素的索引入队。这样保证了队列的单调递减性。
   b. 出队 (Dequeue): 当队首元素的索引已经移出当前窗口范围时，将队首索引出队。
4. 获取窗口最大值：由于队列是单调递减的，队首索引对应的元素始终是当前窗口的最大值。
5. 记录结果：当窗口形成完整大小 (k个元素) 后，开始记录每个窗口的最大值 (即队首元素)。
*/
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)-k+1) // 预先分配好空间
	var q []int
	for i, x := range nums {
		// 1. 入
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1] // 维护 q 的单调性
		}
		q = append(q, i) // 入队
		// 2. 出
		if i-q[0] >= k { // 队首已经离开窗口了
			q = q[1:] // Go 的切片是 O(1) 的
		}
		// 3. 记录答案
		if i >= k-1 {
			// 由于队首到队尾单调递减，所以窗口最大值就是队首
			ans = append(ans, nums[q[0]])
		}
	}
	return ans
}

/*
算法思路:暴力解,把每个窗口的最大值都求出来然添加到答案中
优化:cnt计数表示最大值的个数,如果所有的最大值被移除窗口,才重新计算最大值
如果新增的数字大于最大值,则更新最大值,且重新记cnt=1
如果新增的数字等于最大值,则cnt++
如果移除的数字等于最大值,则cnt--,若cnt<=0,则表明最大值被清空,重新计算最大值
*/
func maxSlidingWindow1(nums []int, k int) []int {
	var res []int
	left, right := 1, k
	num1, cnt := maxNum(nums[0:k])
	res = append(res, num1)
	for right < len(nums) {
		if num1 < nums[right] {
			num1 = nums[right]
			cnt = 1
		} else if num1 == nums[right] {
			cnt++
		} else if num1 == nums[left-1] {
			cnt--
			if cnt <= 0 {
				num1, cnt = maxNum(nums[left : right+1])
			}
		}
		res = append(res, num1)
		left++
		right++
	}
	return res
}
func maxNum(nums []int) (int, int) {
	res := math.MinInt
	cnt := 0
	for _, num := range nums {
		if num > res {
			res = num
			cnt = 1
		} else if res == num {
			cnt++
		}
	}
	return res, cnt
}
