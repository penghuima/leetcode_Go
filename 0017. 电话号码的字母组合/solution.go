package leetcode

func letterCombinations(digits string) []string {
	var res []string
	var path []byte
	//先使用哈希表将数字和字母对应起来
	phoneMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	//如果不写这个判断，当digits 是空字符串 “”时，backTrack还是会在res里追加一个空字符串，变成[""]而不是[]
	if len(digits) == 0 {
		return []string{}
	}
	var backTrack func(int)
	backTrack = func(index int) {
		//如果满足条件
		if index== len(digits) { //或者len(path)
			res = append(res, string(path))
			return
		}
		//获取手机号数字
		digit := digits[index]
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			path = append(path, letters[i])
			backTrack(index + 1) //通过这一步向下驱动
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}
