package DoublePointer

//88. 合并两个有序数组
/*
合并两个按非递减顺序排列的整数数组nums1 和 nums2，使合并后的数组同样按 非递减顺序 排列。
*/
/*
算法思路:双指针
两个指针指向两个数组前面，比较两个指针指向的元素，小的放入新数组，指针后移,需要额外的空间
反向双指针
两个指针分别指向两个数组的末尾，比较两个指针指向的元素，大的放入新数组，指针前移,不需要额外的空间
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
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

func merge1(nums1 []int, m int, nums2 []int, n int) {
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
