package Greedy

//621.任务调度器
/*
算法思路:贪心算法
先统计任务频率,然后找到任务频率最高的任务
把任务频率最高的任务按间隔n隔开
再计算间隔中的空闲无任务格子
如果所有的间隔n不够剩下的任务分,则最小长度就是任务的长度
如果够分,就是任务长度加上最终的剩余格子
*/

func leastInterval(tasks []byte, n int) int {
	// Step 1: 计算任务频率
	taskCount := make(map[byte]int)
	for _, task := range tasks {
		taskCount[task]++
	}

	// Step 2: 找到任务中频率最高的任务次数
	//maxFreq是任务中出现频率最高的任务的出现次数
	maxFreq := 0
	//maxCount是出现频率等于 maxFreq 的任务的数量
	maxCount := 0
	for _, count := range taskCount {
		if count > maxFreq {
			maxFreq = count
			maxCount = 1
		} else if count == maxFreq {
			maxCount++
		}
	}

	// Step 3: 计算最小时间
	// 使用公式：(maxFreq - 1) * (n + 1) + maxCount
	//partCount表示有几个间隔为n的分区
	partCount := maxFreq - 1
	//emptySlots计算假如所有的任务的频率都是最大频率时的空闲槽位数
	emptySlots := partCount * (n - maxCount + 1)
	//availableTasks表示除了出现 maxFreq 次的任务外的任务数量
	availableTasks := len(tasks) - maxFreq*maxCount
	//idleSlots表示最终需要的空闲槽位数
	idleSlots := max(0, emptySlots-availableTasks)

	return len(tasks) + idleSlots
}
