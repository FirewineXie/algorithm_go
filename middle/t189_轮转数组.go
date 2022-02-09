package middle

//

// 使用额外的内存转换
func rotate(nums []int, k int) {
	n := len(nums)
	result := make([]int, len(nums)) //  如果直接相等，为浅拷贝
	copy(result, nums)
	for i := 0; i < n; i++ {
		rotateNum := i + k
		if rotateNum <= n-1 {
			// 如果小于，证明不需要求模，直接进行插入
			nums[rotateNum] = result[i]
		} else {
			temp := rotateNum % n
			nums[temp] = result[i]
		}
	}
}

// 通过遍历交换的形式
// 缺点：  两个遍历，时间较久，，容易超时
func rotate2(nums []int, k int) {
	var temp, previous int
	for i := 0; i < k; i++ {
		previous = nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			temp = nums[j]
			nums[j] = previous
			previous = temp
		}
	}
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// 利用特性反转
func rotate3(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
