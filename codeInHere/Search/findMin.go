package Search

//153. 寻找旋转排序数组中的最小值
//已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。如:[0,1,2,4,5,6,7] 经旋转后可能变为:[4,5,6,7,0,1,2]
/*
算法思路:二分查找
1. 由于数组是原本升序数组经过若干次旋转得到的，所以它可以看作是由两个有序子数组拼接而成。
2. 最小值即为旋转点后的第一个元素。
3. 使用二分查找法，通过比较中间元素与右端元素的大小关系，判断最小值所在的区间。
   - 如果中间元素大于右端元素，说明最小值在右半部分。
   - 如果中间元素小于等于右端元素，说明最小值在左半部分（包括中间元素）。
4. 不断缩小查找区间，直到找到最小值。
*/
func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if mid < n-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[left]
}

// 每次比较中间元素和右端元素：
// - 如果中间元素小于右端元素，说明最小值在左半部分（包括中间元素），high = pivot
// - 否则最小值在右半部分（不包括中间元素），low = pivot + 1
// 最终 low 指向最小值的位置。
func findMin1(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}
	return nums[low]
}
