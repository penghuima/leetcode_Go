package leetcode

func isPartition(s string, start, end int) bool {
	for ; start < end; {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func partition(s string) [][]string {
	var res [][]string
	var path []string //切割字符串集合
	var backTrack func(int)
	backTrack = func(index int) {
		//满足结束条件
		if index == len(s) {
			res = append(res, append([]string(nil), path...))
		}
		for i := index; i < len(s); i++ {
			if isPartition(s, index, i) {
				path = append(path, s[index:i+1])
			}else {
				continue
			}
			backTrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}