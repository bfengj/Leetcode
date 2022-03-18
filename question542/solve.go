package question542


func updateMatrix(mat [][]int) [][]int {

	n,m := len(mat),len(mat[0])
	queue := [][]int{}
	hasDeal := make([][]int,n)
	for k,_ := range hasDeal{
		hasDeal[k] = make([]int,m)
	}
	for i:=0;i<n;i++ {
		for j:=0;j<m;j++ {
			if mat[i][j]==0 {
				queue = append(queue, []int{i,j})
				hasDeal[i][j] = 1
			}
		}
	}
	xx := []int{-1,1,0,0}
	yy := []int{0,0,-1,1}
	for i:=0;i<len(queue);i++ {
		nowElem := queue[i]
		for j:=0;j<4;j++ {
			mx,my := nowElem[0]+xx[j],nowElem[1]+yy[j]
			if mx>=0&&my>=0&&mx<n&&my<m&&hasDeal[mx][my]==0 {
				hasDeal[mx][my] = hasDeal[nowElem[0]][nowElem[1]]+1
				mat[mx][my] = hasDeal[mx][my]-1
				queue = append(queue, []int{mx,my})
			}
		}
	}
	return mat
}