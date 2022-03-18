package question167

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	left,right := 0,n-1

	for left<right {
		sum := numbers[left]+numbers[right]
		if sum==target {
			return []int{left+1,right+1}
		}else if sum<target {
			left++
		}else {
			right--
		}
	}
	return []int{}
}