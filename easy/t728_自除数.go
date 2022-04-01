package easy

import (
	"strconv"
	"strings"
)

func selfDividingNumbers(left int, right int) []int {
	var result []int
	rr := true
	for i := left; i <= right; i++ {
		tmp := strconv.Itoa(i)
		split := strings.Split(tmp, "")

		for j := 0; j < len(split); j++ {
			c, _ := strconv.Atoi(split[j])
			if c == 0 {
				rr = false
				break
			}
			if i%c == 0 {
				continue
			} else {
				rr = false
				break
			}
		}
		if rr {
			result = append(result, i)

		}
		rr = true
	}
	return result
}
