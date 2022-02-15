package middle

// 深度搜索
var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			temp := dfs(grid, i, j)
			if maxArea < temp {
				maxArea = temp
			}
		}
	}
	return maxArea
}

func dfs(grid [][]int, x, y int) int {
	maxArea := 0
	if grid[x][y] == 1 {
		maxArea++
		grid[x][y] = 0
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx >= 0 && mx < len(grid) && my >= 0 && my < len(grid[0]) {
				maxArea += dfs(grid, mx, my)
			}
		}
	}
	return maxArea
}
