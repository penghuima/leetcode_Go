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

/*
func partition(s string) [][]string {
    var tmpString []string//切割字符串集合
    var res [][]string//结果集合
    backTracking(s,tmpString,0,&res)
    return res
}
func backTracking(s string,tmpString []string,startIndex int,res *[][]string){
    if startIndex==len(s){//到达字符串末尾了
        //进行一次切片拷贝，怕之后的操作影响tmpString切片内的值
        t := make([]string, len(tmpString))
		copy(t, tmpString)
        *res=append(*res,t)
    }
    for i:=startIndex;i<len(s);i++{
        //处理（首先通过startIndex和i判断切割的区间，进而判断该区间的字符串是否为回文，若为回文，则加入到tmpString，否则继续后移，找到回文区间）（这里为一层处理）
        if isPartition(s,startIndex,i){
            tmpString=append(tmpString,s[startIndex:i+1])
        }else{
            continue
        }
        //递归
        backTracking(s,tmpString,i+1,res)
        //回溯
        tmpString=tmpString[:len(tmpString)-1]
    }
}
//判断是否为回文
func isPartition(s string,startIndex,end int)bool{
    left:=startIndex
    right:=end
    for ;left<right;{
        if s[left]!=s[right]{
            return false
        }
        //移动左右指针
        left++
        right--
    }
    return true
}
*/
