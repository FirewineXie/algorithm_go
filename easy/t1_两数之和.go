package easy

// 第一种笨办法，，O(n^2) 遍历数组，获取答案
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 核心思想上是，余数补充，，减少target-x  的遍历时间
// 使用额外的空间，减少一次遍历
func twoSum2(nums []int, target int) []int {

	hashTable := map[int]int{}

	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
