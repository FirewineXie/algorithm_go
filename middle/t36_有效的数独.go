package middle

func isValidSudoku(board [][]byte) bool {
	//	记录某行，某位数字是否已经被摆放
	row := [9][9]bool{}
	//	记录某列，某位数字是否已经被摆放
	col := [9][9]bool{}
	// 记录3*3 ,某位数字是否已经被摆放
	block := [9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				blockIndex := i/3*3 + j/3
				if row[i][num] || col[j][num] || block[blockIndex][num] {
					return false
				} else {
					row[i][num] = true
					col[j][num] = true
					block[blockIndex][num] = true
				}
			}
		}
	}
	return true
}
