package Merge

//88.合并两个有序数组
/*
算法思路:
方法一:从nums1末尾填充
从nums1和nums2的有真实数的末尾开始比较,将较大的元素填充到nums1的末尾
方法二:新建一个数组
从nums1和nums2的有真实数的开头开始比较,将较小的元素填充到新数组的前面知道填满
*/

func mergeTwoArray(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	tail := m + n - 1
	var curr int
	for ; p1 >= 0 || p2 >= 0; tail-- {
		if p1 == -1 {
			curr = nums2[p2]
			p2--
		} else if p2 == -1 {
			curr = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			curr = nums1[p1]
			p1--
		} else {
			curr = nums2[p2]
			p2--
		}
		nums1[tail] = curr
	}
}

func mergeTwoArray2(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	nums3 := make([]int, 0, m+n)
	for {
		if i == m {
			nums3 = append(nums3, nums2[j:]...)
			break
		}
		if j == n {
			nums3 = append(nums3, nums1[i:]...)
			break
		}
		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i++
		} else {
			nums3 = append(nums3, nums2[j])
			j++
		}
	}
	copy(nums1, nums3)
}
