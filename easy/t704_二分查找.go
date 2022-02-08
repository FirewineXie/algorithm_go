package easy

// question：给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target
//，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 如果就是想用递归查找呢

func search2(nums []int, target int) int {
	return innerSearch(nums, target, 0)
}

func innerSearch(nums []int, target int, order int) int {
	if len(nums) < 1 || order < 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] != target {
			return -1
		}
	}
	mid := len(nums) / 2
	num := nums[mid]
	if num == target {
		return order + mid
	} else if num > target {
		return innerSearch(nums[0:mid], target, order)
	} else if num < target {
		return innerSearch(nums[mid:], target, order+mid)
	} else {
		return -1
	}
}
