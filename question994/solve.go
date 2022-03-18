package question994

func orangesRotting(grid [][]int) int {
	n,m := len(grid),len(grid[0])

	queue := [][]int{}
	hasDeal := make([][]int,n)
	for k,_ := range hasDeal{
		hasDeal[k] = make([]int,m)
	}
	freshNum := 0
	for i:=0;i<n;i++ {
		for j:=0;j<m;j++ {
			if grid[i][j]==2 {
				queue = append(queue, []int{i,j})
				hasDeal[i][j] = 1
			}else if grid[i][j]==1 {
				freshNum++
			}
		}
	}
	if freshNum==0 {
		return 0
	}

	xx := []int{-1,1,0,0}
	yy := []int{0,0,-1,1}
	var res int
	for i:=0;i<len(queue);i++ {
		nowElem := queue[i]
		for j:=0;j<4;j++ {
			mx,my := nowElem[0]+xx[j],nowElem[1]+yy[j]
			if mx>=0&&my>=0&&mx<n&&my<m&&hasDeal[mx][my]==0&&grid[mx][my]==1 {
				hasDeal[mx][my] = hasDeal[nowElem[0]][nowElem[1]]+1
				res = hasDeal[mx][my]-1
				grid[mx][my] = 2
				freshNum--
				queue = append(queue, []int{mx,my})
			}
		}
	}
	if freshNum!=0 {
		return -1
	}
	return res

}
