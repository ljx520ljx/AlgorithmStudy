package Greedy

//跳跃游戏:
//给你一个非负整数数组nums,你最初位于数组的第一个下标数组中的每个元素代表你在该位置可以跳跃的最大长度。
//判断你是否能够到达最后一个下标,如果可以,返回 true;否则返回false
/*
算法思路:贪心算法
初始化一个变量 maxReach 为 0，表示我们当前能到达的最远位置。
遍历数组中的每个位置 i，并检查：
如果当前位置 i 超出了 maxReach，说明我们无法到达该位置，返回 false。
更新 maxReach 为 max(maxReach, i + nums[i])，即考虑当前位置 i 的跳跃能力。
如果 maxReach 已经大于或等于数组的最后一个位置，返回 true。
遍历结束后，返回 false，说明无法到达最后一个位置。
*/

func canJump(nums []int) bool {
	maxReach := 0 // 最远能到达的位置
	n := len(nums)

	// 遍历数组中的每个位置
	for i := 0; i < n; i++ {
		if i > maxReach {
			// 如果当前位置超出了最远可达位置，说明无法到达该位置
			return false
		}

		// 更新最远可达位置
		maxReach = max(maxReach, i+nums[i])

		// 如果最远可达位置已经超过或等于最后一个位置，返回 true
		if maxReach >= n-1 {
			return true
		}
	}
	// 默认返回 false
	return false
}

// 简洁写法
func canJump1(nums []int) bool {
	mx := 0
	for i, jump := range nums {
		if i > mx { // 无法到达 i
			return false
		}
		mx = max(mx, i+jump)   // 从 i 最右可以跳到 i + jump
		if mx >= len(nums)-1 { // 可以跳到 n-1
			break
		}
	}
	return true
}
