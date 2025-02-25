package SpecialAlgorithm

//189. 轮转数组
//给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
/*
算法思路:
先简化轮转次数,只需要轮转k%n次,再考虑清楚轮转k次实际是将数组的最后k个元素翻转到数组的最前面
所以先翻转整个数组,然后翻转前k个元素,然后翻转k到n-1个元素
其实这道题考的就是每个元素轮转之后的位置,然后将每个元素放到该位置上
*/

func rotatearray(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
func reverse(nums []int) {
	i, n := 0, len(nums)
	for i < n/2 {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
		i++
	}
}
