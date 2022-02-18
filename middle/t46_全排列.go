package middle

/*
@Time : 2022/2/18 23:38
@Author : Firewine
@File : t46_全排列
@Software: GoLand
@Description:
*/

func permute(nums []int) [][]int {

	return dfsPermute(len(nums), 0, nums)
}

func dfsPermute(n, first int, nums []int) (result [][]int) {

	if first == n {
		var tmp []int
		tmp = append(tmp, nums...)
		result = append(result, tmp)
		return result
	}
	// 横向遍历
	for i := first; i < n; i++ {
		// 交换
		swap := nums[first]
		nums[first] = nums[i]
		nums[i] = swap

		result = append(result, dfsPermute(n, first+1, nums)...)
		swap = nums[i]
		nums[i] = nums[first]
		nums[first] = swap

	}
	return
}
