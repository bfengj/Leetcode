package question695

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	xx := []int{-1, 0, 0, 1}
	yy := []int{0, -1, 1, 0}
	for i := 0;i<len(grid);i++ {
		for j := 0;j<len(grid[i]);j++ {
			if grid[i][j]==0 {
				continue
			}
			nowArea := 1
			queue := [][]int{}
			queue = append(queue,[]int{i,j})
			grid[i][j] = 0
			for k := 0; k < len(queue);k++ {
				nowEle := queue[k]
				for p := 0; p < 4; p++ {
					mx, my := nowEle[0]+xx[p], nowEle[1]+yy[p]
					if mx >= 0 && my >= 0 && mx < len(grid) && my < len(grid[0]) && grid[mx][my] != 0 {
						grid[mx][my] = 0
						nowArea++
						queue = append(queue, []int{mx, my})
					}
				}
			}
			if nowArea>res {
				res = nowArea
			}
		}
	}
	return res
}