package middle

import "fmt"

// 使用m+n 的空间

func setZeroes(matrix [][]int) {
	row := map[int]bool{} // 标记那些列需要置空
	col := map[int]bool{} // 标记那些行需要置空

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {

			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true

			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if _, ok := row[i]; ok {
				matrix[i][j] = 0
			}
			if _, ok := col[j]; ok {
				matrix[i][j] = 0
			}
		}
	}
	fmt.Println(&matrix)
}

func setZeroes2(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}

}
