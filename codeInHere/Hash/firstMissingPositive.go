package Hash

//41. 缺失的第一个正数
/*
算法思路:原地哈希
将数组中所有小于等于 0 的数修改为 N+1；
再把数组当作哈希表,key值就是数组下表,把val值标记为负数,
最后遍历数组,第一个正数的下标+1就是缺失的第一个正数
注意:中间因为标记的时候把数组元素标记为了负数,所以再遍历数组时要还原成正数
*/
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for _, num := range nums {
		num = abs(num) - 1
		if num < n {
			nums[num] = -abs(nums[num])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
