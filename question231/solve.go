package question231


func isPowerOfTwo(n int) bool {
	return (n>0)&&(n&(n-1)==0)
}
/*func isPowerOfTwo(n int) bool {
	if n==1 {
		return true
	}
	num := 1
	for num<=n {
		num = num<<1
	}
	if num==n {
		return true
	}
	return false
}*/