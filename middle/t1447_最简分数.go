package middle

import "fmt"

// 缺点是，只能通过暴力，求解最大公约数等于1
func simplifiedFractions(n int) []string {
	var result []string
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {

				result = append(result, fmt.Sprintf("%d/%d", j, i))
			}
		}
	}

	return result
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
