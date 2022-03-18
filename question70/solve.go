package question70


/*var cache []int = make([]int,100)
func climbStairs(n int) int {
	if cache[n-1]!=0 {
		return cache[n-1]
	}
	if n==1||n==2 {
		cache[n-1] = n
		return n
	}
	cache[n-1] = climbStairs(n-1)+climbStairs(n-2)
	return cache[n-1]
}*/

func climbStairs(n int) int {
	x,y,z :=0,0,1
	for i:=1;i<=n;i++ {
		x = y
		y = z
		z = x+y
	}
	return z
}