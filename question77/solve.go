package question77

var res [][]int

func combine(n int, k int) [][]int {
	var path []int
	res = [][]int{}
	backtrace(path, n, k, 1)
	return res
}
func backtrace(path []int, n int, k int, start int) {
	if len(path) == k {
		//后续修改path，会影响到res中的内容。
		temp := make([]int, k)
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := start; i <= n; i++ {
		//剪枝
		if i<=1+n-k+len(path) {
			path = append(path, i)
			backtrace(path, n, k, i+1)
			path = path[:len(path)-1]
		}
	}
}
