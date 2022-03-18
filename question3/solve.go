package question3

func max(x,y int) int{
	if x>y{
		return x
	}else{
		return y
	}
}
/*func lengthOfLongestSubstring(s string) int {
	n := len(s)
	mmax := max(n,128)
	maxLength := 0
	left,right := 0,0
	for ;left<n;left++ {
		myMap := make(map[byte]int)
		for right=left;right<n&&right-left<mmax;right++ {
			if myMap[s[right]]>0 {
				break
			}
			myMap[s[right]]++
		}
		maxLength = max(maxLength,right-left)
		if right==n {
			return maxLength
		}
		left += strings.Index(string([]byte(s)[left:right]), string(s[right]))
		fmt.Println(left)
	}
	return maxLength
}*/
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	mmax := max(n,128)
	maxLength := 0
	myMap := make(map[byte]int)
	left,right := 0,0
	for ;left<n;left++ {
		if left != 0 {
			myMap[s[left-1]]--
		}
		for ;right<n&&right-left<mmax;right++ {
			if myMap[s[right]]>0 {
				break
			}
			myMap[s[right]]++
		}
		maxLength = max(maxLength,right-left)
	}
	return maxLength
}