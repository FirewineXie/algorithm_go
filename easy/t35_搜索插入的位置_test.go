package easy

import (
	"fmt"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	//nums := []int{2, 5}
	target := 0
	fmt.Println(searchInsert(nums, target))

}
