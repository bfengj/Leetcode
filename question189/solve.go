package question189


func rotate(nums []int, k int)  {
	n := len(nums)
	k = k%n
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int){
	n := len(nums)
	for i:=0;i<n/2;i++{
		temp := nums[i]
		nums[i]=nums[n-i-1]
		nums[n-i-1] = temp
	}
}