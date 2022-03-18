package question198

func max(x,y int) int {
	if x>y {
		return x
	}
	return y
}
func rob(nums []int) int {
	n := len(nums)
	if n==1 {
		return nums[0]
	}
	first,second := nums[0],max(nums[0],nums[1])

	for i:=3;i<=n;i++ {
		first,second = second,max(second,nums[i-1]+first)
	}
	return second
}

/*func rob(nums []int) int {
	n := len(nums)
	res := make([]int,n)
	for i:=1;i<=n;i++ {
		if i==1 {
			res[i-1] = nums[i-1]
		}else if i==2 {
			res[i-1] = max(nums[i-1],nums[i-2])
		}else{
			res[i-1] = max(res[i-2],res[i-3]+nums[i-1])
		}
	}
	return res[n-1]
}
*/