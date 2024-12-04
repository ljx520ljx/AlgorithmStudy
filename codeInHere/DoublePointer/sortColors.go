package DoublePointer

//颜色分类
//给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//必须在不使用库内置的 sort 函数的情况下解决这个问题。
/*
算法思路:交换位置
用nums0记录0的数量,nums2记录2的数量
碰到0,把它与数组前面的元素交换位置,nums[nums0-1], nums[i] = nums[i], nums[nums0-1]
碰到2,把它与数组后面的元素交换位置,nums[n-nums2], nums[i] = nums[i], nums[n-nums2]
注意:由于遍历到后面会重复遍历到已经被我们放在后面的2,所以nums2会超出实际的2的数量
所以当n-nums2 <= i时,i为当前索引,要break禁止交换位置,否则会和前面的1交换位置
*/

func sortColors(nums []int) {
	n := len(nums)
	nums0 := 0
	nums2 := 0
	for i := 0; i < n; i++ {
		for nums[i] == 0 || nums[i] == 2 {
			if nums[i] == 0 {
				nums0++
				if nums0-1 == i {
					break
				}
				nums[nums0-1], nums[i] = nums[i], nums[nums0-1]
			}
			if nums[i] == 2 {
				nums2++
				if n-nums2 <= i {
					break
				}
				nums[n-nums2], nums[i] = nums[i], nums[n-nums2]
			}
		}
	}
}

//方法二是双指针
/*
p0指向0区后的第一个数,p2指向2区前的第一个数
*/
func sortColors1(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		//这里需要for重复检验,因为被交换的元素也有可能是2
		for i <= p2 && nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
