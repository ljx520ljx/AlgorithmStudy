package DoublePointer

//两数之和
/*
算法思路::用哈希表把时间复杂度降为O(n)
也可以双层遍历时间复杂度为O(n^2)
*/

func twoSum(nums []int, target int) []int {
	// 初始化一个哈希表，存储数字和对应的索引
	hashMap := make(map[int]int)

	// 遍历数组
	for i, num := range nums {
		// 计算需要查找的另一个数字
		complement := target - num

		// 检查哈希表中是否存在这个数字
		if index, ok := hashMap[complement]; ok {
			return []int{index, i} // 找到两个数，返回它们的索引
		}

		// 将当前数字及其索引加入哈希表
		hashMap[num] = i
	}

	// 如果没有找到结果，返回空数组
	return []int{}
}
