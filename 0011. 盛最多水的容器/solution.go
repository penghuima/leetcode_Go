package leetcode

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	areaMax := 0
	for i < j {
		width := j - i
		if height[i] > height[j] {
			areaMax = max(areaMax, height[j]*width)
			j--
		} else {
			areaMax = max(areaMax, height[i]*width)
			i++
		}
	}
	return areaMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
