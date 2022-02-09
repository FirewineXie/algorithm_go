package easy

import (
	"math"
)

// 先使用粗暴的方法解决问题

func countKDifference(nums []int, k int) int {

	num := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if int(math.Abs(float64(nums[i]-nums[j]))) == k {
				num++
			}
		}
	}
	return num
}

// 1 <= nums[i] <= 100
//满足 i < j 且 |nums[i] - nums[j]| == k 。
// 上述条件，降低了这道题的难度

// 优化解决方法 : 存在问题： map 只能验证一个，无法验证多个
// 使用空间置换时间
func countKDifference2(nums []int, k int) int {
	remainder := map[int]int{}
	var ans int
	for _, num := range nums {
		ans += remainder[num-k] + remainder[num+k]
		remainder[num]++
	}
	return ans

}
