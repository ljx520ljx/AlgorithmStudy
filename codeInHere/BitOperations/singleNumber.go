package BitOperations

//只出现一次的数字,找到数组里只出现了一次的数字,要求时间复杂度0(n),空间复杂度0(1)
/*
算法思路:排序或者位运算
排序后nums[i]!=nums[i+1]:return nums[i];i==len(nums)-1:return nums[i]
利用异或运算 a⊕a=0 的性质，我们可以用异或来「消除」所有出现了两次的元素，最后剩下的一定是只出现一次的元素。
*/

import "slices"

// 排序,时间复杂度O(nlog n)
func singleNumber(nums []int) int {
	slices.Sort(nums)
	for i := 0; i < len(nums); i = i + 2 {
		if i == len(nums)-1 {
			return nums[i]
		} else if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

// 异或位运算,时间复杂度0(n)
func singleNumber1(nums []int) (ans int) {
	for _, x := range nums {
		ans ^= x
	}
	return
}
