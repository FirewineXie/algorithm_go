package middle

import "sort"

//给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足
//“对于每个 0 <= i < len(arr) / 2，都有 arr[2 * i + 1] = 2 * arr[2 * i]” 时，返回 true；否则，返回 false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/array-of-doubled-pairs
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 1. 暴力解法

func canReorderDoubled(arr []int) bool {

	cnt := make(map[int]int, len(arr))

	for _, value := range arr {
		cnt[value]++
	}

	if cnt[0]%2 == 1 {
		return false
	}

	vals := make([]int, 0, len(cnt))
	for x := range cnt {
		vals = append(vals, x)
	}
	sort.Slice(vals, func(i, j int) bool { return abs(vals[i]) < abs(vals[j]) })

	for _, x := range vals {
		if cnt[2*x] < cnt[x] { // 无法找到足够的 2x 与 x 配对
			return false
		}
		cnt[2*x] -= cnt[x]
	}
	return true

}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
