package easy

import "sort"

//给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。

// 1. 使用hash表，利于hash 达到唯一值判断，，但是存在遍历全局的情况
func containsDuplicate(nums []int) bool {

	numsMap := make(map[int]bool, len(nums))

	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			return true
		} else {
			numsMap[num] = true
		}
	}
	return false
}

// 1. 排序, 内部实现为quickSort 快排
func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
