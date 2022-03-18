package question120

func min(x,y int)int  {
	if x>y {
		return y
	}
	return x
}
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, i+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			if i==0 {
				res[i][j] = triangle[i][j]
			}else if j==0 {
				res[i][j] = res[i-1][j]+triangle[i][j]
			}else if j==i{
				res[i][j] = res[i-1][j-1]+triangle[i][j]
			}else{
				res[i][j] = min(res[i-1][j],res[i-1][j-1])+triangle[i][j]
			}
		}
	}
	minElem := res[n-1][0]
	for i := 1; i <n ; i++ {
		if res[n-1][i]<minElem {
			minElem = res[n-1][i]
		}
	}
	return minElem
}

