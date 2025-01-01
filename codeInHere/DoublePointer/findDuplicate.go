package DoublePointer

//287. 寻找重复数
/*
算法思路:快慢指针
快慢指针:将数组元素视为链表节点的 next 指针，slow 指针每次走一步，即 slow = nums[slow]，
fast 指针每次走两步，即 fast = nums[nums[fast]]，通过快慢指针的移动最终会在环内相遇。
环入口:当 slow 和 fast 相遇后，将 slow 重新设置为 nums[0]，然后 slow 和 fast 以相同速度移动，
它们会在环入口相遇，该环入口就是重复元素所在位置，
这是通过数学推导 2(m + x) = m + x  + lc 得出 m = lc - x 的结论，
所以当 slow 走 m 步到达环入口时，fast 也会到达环入口。
*/
//暴力解
func findDuplicate(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return 0
}

// 快慢指针
func findDuplicate1(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
