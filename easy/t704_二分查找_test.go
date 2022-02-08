package easy

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	//nums := []int{2, 5}
	target := 2
	fmt.Println(search2(nums, target))

}
