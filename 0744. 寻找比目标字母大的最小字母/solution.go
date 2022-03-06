package leetcode

//这个题不就是找到 target 字符出现的最右边的位置    然后返回该位置的后一个字符就好了
func nextGreatestLetter(letters []byte, target byte) byte {
	//题目已经给了长度大于2
	index := searchTargetRight(letters, target)
	if index == len(letters)-1 {
		return letters[0]
	}
	return letters[index+1]

}
func searchTargetRight(letters []byte, target byte) int {
	left, right := 0, len(letters)
	for left < right {
		mid := left + (right-left)>>1
		if letters[mid] > target {
			right = mid
		} else {
			left = mid + 1 //即使找到了相等 也将left +1  最后返回的时候返回 left-1即可
		}
	}
	//一定是存在的这个字符
	return left - 1
}
