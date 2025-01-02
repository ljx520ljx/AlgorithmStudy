package SpecialAlgorithm

func leastInterval(tasks []byte, n int) int {
	// Step 1: 计算任务频率
	taskCount := make(map[byte]int)
	for _, task := range tasks {
		taskCount[task]++
	}

	// Step 2: 找到任务中频率最高的任务次数
	maxFreq := 0
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
	partCount := maxFreq - 1
	emptySlots := partCount * (n - maxCount + 1)
	availableTasks := len(tasks) - maxFreq*maxCount
	idleSlots := max(0, emptySlots-availableTasks)

	return len(tasks) + idleSlots
}
