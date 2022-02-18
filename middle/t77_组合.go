package middle

/*
@Time : 2022/2/18 22:08
@Author : Firewine
@File : t77_组合
@Software: GoLand
@Description:
*/

// 回溯： 经典问题，皇后棋，，回溯就是暴力组合的优雅展示
func combine(n int, k int) [][]int {
	var result [][]int

	if k < 1 || k > n {
		return result
	}
	var t []int
	result = dfsCombine(1, n, k, t)
	return result
}

func dfsCombine(begin int, n int, k int, t []int) (result [][]int) {

	// 临界条件
	if k == len(t) {
		var tmp []int
		tmp = append(tmp, t...)
		result = append(result, tmp)
		return result
	}

	for i := begin; i <= n; i++ {
		t = append(t, i)
		result = append(result, dfsCombine(i+1, n, k, t)...)
		t = t[0 : len(t)-1]
	}
	return
}
