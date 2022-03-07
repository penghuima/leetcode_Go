package leetcode

//返回的结果肯定是 k 个连续的 利用二分法找这段数组的左边界
func findClosestElements(arr []int, k int, x int) []int {
	res := []int{}
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)>>1
		//防止越界 mid+k
		if mid+k < len(arr) && x-arr[mid] > arr[mid+k]-x { //右边界比左边界更靠近x 向右寻找左边界
			left = mid + 1
		} else {
			right = mid
		}
	}
	//结果存在
	res = append(res, arr[left:left+k]...)
	return res
}
