package question46

import "fmt"

var res [][]int
func permute(nums []int) [][]int {
	path := []int{}
	res = [][]int{}
	backtrace(path,nums)
	return res
}
func isInSlice(path []int,value int) bool{
	for _,v:=range path{
		if v == value {
			return true
		}
	}
	return false
}
func backtrace(path []int,nums []int) {
	if len(path) == len(nums) {
		fmt.Println(len(nums))
		//后续修改path，会影响到res中的内容。
		temp := make([]int, len(nums))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := 0; i <len(nums); i++ {
		if !isInSlice(path,nums[i]) {
			fmt.Println(nums[i])
			path = append(path, nums[i])
			backtrace(path,nums)
			path = path[:len(path)-1]
		}
	}
}
