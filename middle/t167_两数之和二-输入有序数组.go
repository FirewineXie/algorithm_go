package middle

func twoSum(numbers []int, target int) []int {

	remianders := map[int]int{}

	for i := 0; i < len(numbers); i++ {
		if n, ok := remianders[target-numbers[i]]; ok {
			return []int{n + 1, i + 1}
		}
		remianders[numbers[i]] = i
	}
	return nil
}

func twoSum2(numbers []int, target int) []int {

	left, right := 0, len(numbers)-1

	for true {
		if numbers[left]+numbers[right] == target {
			break
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{left + 1, right + 1}
}

// 第三种二分搜索，固定一个，搜索另外一个，
