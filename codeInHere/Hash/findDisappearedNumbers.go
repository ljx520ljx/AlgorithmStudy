package Hash

//找出所有数组中消失的数字
//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。
//请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。
/*
算法思路:
方法一:可以用hash表存数组中所有的数字,一但出现一个数不在数组中但是属于1-n,
就把这个数添加到答案数组中,时间复杂度是O(N),空间复杂度是O(N)
方法二:实现原地修改
遍历 nums，每遇到一个数 x，就让 nums[x−1] 增加 n。由于 nums 中所有数均在 [1,n] 中，
增加以后，这些数必然大于 n。最后我们遍历 nums，若 nums[i] 未大于 n，就说明没有遇到过数 i+1。这样我们就找到了缺失的数字。
*/

func findDisappearedNumbers(nums []int) []int {
	var res []int
	numMap := make(map[int]bool, len(nums))
	for _, j := range nums {
		numMap[j] = true
	}
	for i := 1; i <= len(nums); i++ {
		if !numMap[i] {
			res = append(res, i)
		}
	}
	return res
}

func findDisappearedNumbers1(nums []int) (ans []int) {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return
}
