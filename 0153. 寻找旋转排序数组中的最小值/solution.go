package leetcode


func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] { //中间值小于最后一个数，说明最小值在左侧
			high = pivot
		} else if nums[pivot] > nums[high] {
			low = pivot + 1
		} else { //相等
			low = pivot + 1
		}
	}
	return nums[low]
}