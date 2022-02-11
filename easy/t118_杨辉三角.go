package easy

func generate(numRows int) [][]int {

	result := [][]int{{1}}
	if numRows == 1 {
		return result
	}
	// f(n)[0] = 1
	// f(n)[1] = f(n-1)[0] + f(n-1)[1]
	// f(n)[n-1] = 1
	for i := 2; i <= numRows; i++ {
		var temp []int //长度i+1
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				temp = append(temp, 1)
			} else {
				t := i - 1
				temp = append(temp, result[t-1][j-1]+result[t-1][j])
			}
		}
		result = append(result, temp)

	}
	return result
}
