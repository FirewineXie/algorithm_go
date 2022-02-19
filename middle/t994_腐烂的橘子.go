package middle

var (
	dr = []int{-1, 0, 1, 0}
	dc = []int{0, -1, 0, 1}
)

// 怀橘子不止一个 ，，所以  多源广度优先搜索
func orangesRotting(grid [][]int) int {
	r, c := len(grid), len(grid[0])

	var queue []int
	depth := make(map[int]int)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 2 {
				code := i*c + j
				queue = append(queue, code)
				depth[code] = 0
			}
		}
	}
	ans := 0
	for len(queue) > 0 {
		code := queue[0]
		i := code / c
		j := code % c
		for k := 0; k < 4; k++ {
			nr := i + dr[k]
			nc := j + dc[k]

			if nr >= 0 && nr < r && nc >= 0 && nc < c && grid[nr][nc] == 1 {
				grid[nr][nc] = 2
				ncode := nr*c + nc
				queue = append(queue, ncode)
				depth[ncode] = depth[code] + 1
				ans = depth[ncode]
			}
		}
	}
	for _, row := range grid {
		for _, v := range row {
			if v == 1 {
				return -1
			}
		}
	}
	return ans
}
