package question733

//BFS
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	target := image[sr][sc]
	if target == newColor {
		return image
	}
	image[sr][sc] = newColor
	queue := [][]int{}
	xx := []int{-1, 0, 0, 1}
	yy := []int{0, -1, 1, 0}
	queue = append(queue, []int{sr, sc})
	for i := 0; i < len(queue); i++ {
		nowEle := queue[i]
		for j := 0; j < 4; j++ {
			mx, my := nowEle[0]+xx[j], nowEle[1]+yy[j]
			if mx >= 0 && my >= 0 && mx < len(image) && my < len(image[0]) && image[mx][my] == target {
				image[mx][my] = newColor
				queue = append(queue, []int{mx, my})
			}
		}
	}
	return image
}
