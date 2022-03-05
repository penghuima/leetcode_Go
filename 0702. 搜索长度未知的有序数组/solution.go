package leetcode

import "math"

var nums []int

type ArrayReader struct{}

func (a *ArrayReader) get(index int) int {
	if index < len(nums) {
		return nums[index]
	}
	return math.MaxInt32
}

//上述一个结构体和一个方法是为了防止报错
func search(reader ArrayReader, target int) int {
	if reader.get(0) == target {
		return 0
	}
	left, right := 0, 1
	//确定搜索边界
	for reader.get(right) <= target {  //防止相等的时候  right没有扩容
		left = right
		right *= 2
	}
	for left < right {
		mid := left + (right-left)>>1
		if reader.get(mid) < target {
			left = left + 1
		} else if reader.get(mid) > target {
			right = mid
		} else {
			return mid
		}
	}
	return -1
}
