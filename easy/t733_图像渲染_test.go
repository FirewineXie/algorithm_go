package easy

import (
	"fmt"
	"testing"
)

func Test_floodFill(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	newColor := 2
	fill := floodFill(image, sr, sc, newColor)
	fmt.Println(fill)
}
