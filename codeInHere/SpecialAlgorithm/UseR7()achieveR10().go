package SpecialAlgorithm

//用rand7()实现rand10()

import "math/rand"

func rand7() int {
	return rand.Intn(7) + 1
}

//自己写的思路,凑10个rand7,然后对10取余数+1
//func rand10() int {
//	var result int
//	result = (2*rand7()+rand7()+3*rand7()+4*rand7())%10 + 1
//	return result
//}

//GPT思路

func rand10() int {
	for {
		// 生成 [1,49] 范围的随机数
		num := (rand7()-1)*7 + rand7()
		// 如果 num 在 1 到 40 范围内，则映射到 1 到 10
		if num <= 40 {
			return num%10 + 1
		}
		// 否则重新生成
	}
}
