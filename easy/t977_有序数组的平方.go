package easy

// 1. 非递减顺序，那么就是递增顺序
// 2. 存在正负值 ，那么 平方后， 是两头大，中间下，
// 3. 指针前后相依判断插入数组中

// 优化为，将result 为了 长度为nums 长度固定数组，进行数据填充，这样不需要扩展数组长度，和反转
func sortedSquares(nums []int) []int {
	var result []int
	// 确定1 个边界条件，数组长度为1
	if len(nums) == 1 {
		result = append(result, nums[0]*nums[0])
		return result
	}
	length := len(nums) - 1
	left, right := 0, length

	for true {
		leftS := nums[left] * nums[left]
		rightS := nums[right] * nums[right]
		if leftS >= rightS {
			result = append(result, leftS)
			left++
		} else {
			result = append(result, rightS)
			right--
		}

		if left == right {
			result = append(result, nums[right]*nums[right])
			break
		}
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func sortedSquared2(nums []int) []int {
	result := make([]int, len(nums))
	left, right := 0, len(nums)-1

	for pos := right; pos >= 0; pos-- {
		if v, w := nums[left]*nums[left], nums[right]*nums[right]; v > w {
			result[pos] = v
			left--
		} else {
			result[pos] = w
			right--

		}
	}

	return result
}
