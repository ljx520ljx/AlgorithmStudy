package SpecialAlgorithm

/*多数元素
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/
/*
算法思路:hashmap或者投票法
用hashmap记录每个数出现的次数,最后次数>n/2的就是答案.
投票法:遍历遇到同样的数count++,count==0更换res,否则count--
*/
//hash法
func majorityElement(nums []int) int {
	n := len(nums)
	numMap := make(map[int]int, n)
	for i := range nums {
		numMap[nums[i]]++
	}
	for i := range nums {
		if numMap[nums[i]] > n/2 {
			return nums[i]
		}
	}
	return 0
}

// 投票法
func majorityElement1(nums []int) int {
	counts := 0
	var res int
	for _, j := range nums {
		if j == res {
			counts++
		} else if counts == 0 {
			res = j
			counts++
		} else {
			counts--
		}
	}
	return res
}
