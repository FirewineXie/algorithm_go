package easy

import (
	"fmt"
	"testing"
)

func Test_reverseString(t *testing.T) {
	s := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s)
	for _, b := range s {
		fmt.Printf("%s ", string(b))
	}
	fmt.Println()
}
