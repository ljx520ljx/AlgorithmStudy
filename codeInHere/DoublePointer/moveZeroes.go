package DoublePointer

//283. 移动零
/*
算法思路:双指针
我的思路:先遍历所有数组,用数组index记录保存所有0的下标.
然后遍历index数组,将0后面的每一位元素向前移动一位,直到最后一个0.
因为中途移动了元素,也移动了元素0的位置,所以需要cnt作为位偏移量和index组合得到最新的0的位置.
双指针思路:
i0作为0元素的左边界,i作为0元素的右边界,i0左侧的元素都是非0元素.中间[i0,i)之间的元素都是0元素.
往后遍历数组,当遇到0元素时,就跳过,直到遇到非0元素,就和i0交换位置,然后i0++,这样就保证了i0左侧的元素都是非0元素.
中途0元素池是动态增加的,直到遍历完整个数组,所有的0元素都被移动到了数组的最后面.
*/
//我的思路,时间复杂度O(n),空间复杂度O(n),实际运行时间长很多
func moveZeroes(nums []int) {
	n := len(nums)
	var index []int
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			index = append(index, i)
		}
	}
	zeroLen := len(index)
	cnt := 0
	for i := 0; i < zeroLen; i++ {
		for j := index[i] - cnt; j < n-1; j++ {
			nums[j] = nums[j+1]
		}
		cnt++
	}
	for i := n - zeroLen; i < n; i++ {
		nums[i] = 0
	}
}

// 双指针思路,时间复杂度O(n),空间复杂度O(1)
func moveZeroes1(nums []int) {
	i0 := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[i0] = nums[i0], nums[i]
			i0++
		}
	}
}
