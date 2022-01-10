package leetcode

func allPathsSourceTarget(graph [][]int) (res [][]int) {
	length := len(graph)
	var path []int
	path = append(path, 0)
	var backTrack func(index int) //此时这个index的意思是节点序号的意思
	backTrack = func(index int) {
		//满足条件
		if index == length-1 {
			res = append(res, append([]int(nil), path...))
			return
		}
		//单层逻辑搜素  遍历所有 index 节点可以到达的节点
		for i := 0; i < len(graph[index]); i++ {
			path = append(path, graph[index][i])
			backTrack(graph[index][i])
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return
}
