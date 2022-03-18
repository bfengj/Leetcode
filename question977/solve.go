package question977

func sortedSquares(nums []int) []int {
	var index int = len(nums)-1
	res := make([]int,len(nums))
	left,right := 0,len(nums)-1
	for left<=right {
		if nums[left]*nums[left] >= nums[right]*nums[right] {
			res[index] = nums[left]*nums[left]
			left++
		}else{
			res[index] = nums[right]*nums[right]
			right--
		}
		index--
	}
	return res
}