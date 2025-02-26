package SpecialAlgorithm

import "math/rand"

type RandomizedSet struct {
	nums   []int
	valMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valMap[val]; ok {
		return false
	}
	this.valMap[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.valMap[val]
	if !ok {
		return false
	}
	length := len(this.nums)
	this.nums[index] = this.nums[length-1]
	this.valMap[this.nums[index]] = index
	this.nums = this.nums[:length-1]
	delete(this.valMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
