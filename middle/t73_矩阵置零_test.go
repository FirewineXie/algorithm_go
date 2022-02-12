package middle

import (
	"fmt"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	fmt.Println(&matrix)
	setZeroes(matrix)
	fmt.Println(&matrix)
	fmt.Println(matrix)
}
