package leetcode

import "strconv"

//判断字串是否合理 符合 [0,255] 且数字不能含有前导0
func isNormalIp(s string, start, end int) bool {
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}
	//将字符串转变为 int 如 ”255“-> 255
	checkInt, _ := strconv.Atoi(s[start : end+1])
	if checkInt > 255 {
		return false
	}
	return true
}

//需要优化，多递归了
func restoreIpAddresses(s string) []string {
	var res []string
	var path []string
	var backTrack func(int)
	backTrack = func(index int) {
		if index == len(s) && len(path) == 4 {
			temp := path[0] + "." + path[1] + "." + path[2] + "." + path[3] //拼成1个ip字符串
			res = append(res, temp)
		}
		for i := index; i < len(s); i++ {
			//i-index+1 <= 3这个条件不满足就不要执行 isNormalIp函数了
			if i-index+1 <= 3 && len(path) < 4 && isNormalIp(s, index, i) {
				path = append(path, s[index:i+1])
				backTrack(i + 1)
			} else {
				return
			}
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}
