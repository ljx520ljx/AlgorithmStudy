package Search

//二分查找
/*
算法思路:二分查找
3种情况
mid := left + (right-left)/2
如果nums[mid] == target,return mid
如果nums[mid] < target,left = mid + 1
如果nums[mid] > target,right = mid - 1
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
