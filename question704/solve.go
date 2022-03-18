package question704

func Search(nums []int, target int) int {
	start,end := 0,len(nums)-1
	for start <= end {
		mid := (start+end)/2
		if nums[mid] == target {
			return mid
		}else if nums[mid] < target {
			start = mid+1
		}else {
			end = mid-1
		}
	}
	return -1
}