package easy

import (
	"math"
)

// 动态规划 或者是分治   ：  记得相关概念是，是将 大的任务分成若干个小任务，动态规划 的每个小任务都是有关系的

// 动态规划
func maxSubArray(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]

		}
	}
	return max
}

// 动态规划，还是以上面为主
func maxSubArray2(nums []int) int {
	sum := 0
	max := nums[0]

	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		max = int(math.Max(float64(max), float64(sum)))

	}
	return max

}
