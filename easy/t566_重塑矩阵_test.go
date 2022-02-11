package easy

import (
	"fmt"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	mate := [][]int{{1, 2}, {3, 4}}
	fmt.Println(matrixReshape(mate, 2, 4))

}
