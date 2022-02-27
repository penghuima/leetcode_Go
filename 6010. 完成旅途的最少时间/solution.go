package leetcode

func minimumTime(time []int, totalTrips int) int64 {
	var res int64 = 0
	minTime := time[0]

	for _, v := range time {
		minTime = min(minTime, v)
	}
	var left, right, mid int64 = 0, int64(minTime * totalTrips), 0
	for left < right { // [ ) 区间
		var count int64 = 0
		mid = left + (right-left)>>1
		for _, v := range time {
			count += int64(mid / int64(v))
		}
		if count >= int64(totalTrips) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	//结束条件是 left=right
	res = left
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
