package easy

func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}

	left, right := 0, 1

	for true {
		if right == len(nums) {
			break
		}
		if nums[left] == 0 && nums[right] != 0 {
			nums[left] = nums[right]
			nums[right] = 0
			left++
			right++
			continue
		}

		if nums[left] == 0 && nums[right] == 0 {
			right++
			continue
		}
		if nums[left] != 0 {
			left++
			right++
			continue
		}

	}
	return
}
