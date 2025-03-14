package Search

//33. 搜索旋转排序数组
//整数数组 nums 按升序排列，数组中的值 互不相同 。数组在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为[4,5,6,7,0,1,2] ）。
/*
算法思路:二分查找的变异
1.找到中点
2.判断中点是否等于target,如果等于,则返回中点
3.判断左半部分有序还是右半部分有序
4.判断target是否在左半部分或者右半部分
5.如果在左半部分,则在左半部分查找,否则在右半部分查找
6.如果没找到,则返回-1
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		//二分查找中点
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		//判断左半部分有序还是右半部分有序
		if nums[left] <= nums[mid] {
			//判断target是否在左半部分
			if target >= nums[left] && target < nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			//判断target是否在右半部分
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return -1
}
