package leetcode

//找峰值
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		//上坡必有顶
		if mid+1 < len(nums) && nums[mid] < nums[mid+1] {
			left = left + 1 //寻找最左边界
		} else {
			right = mid
		}
	}
	return left
}
