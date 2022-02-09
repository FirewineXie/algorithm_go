package easy

import "sort"

// 采用双指针的形式，条件满足，然后补充余下数字就可以了
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 右边可能为空
	if n == 0 {
		return
	}
	left, right := m-1, n-1
	order := len(nums1) - 1
	// 左边可能为空
	if m != 0 {

		for true {
			if nums1[left] > nums2[right] {
				nums1[order] = nums1[left]
				left--
				order--
			} else {
				nums1[order] = nums2[right]
				right--
				order--
			}

			if left == -1 || right == -1 {
				// 任何一方先走完了，跳出
				break
			}
		}
	}

	if right != -1 {
		for i := order; i >= 0; i-- {
			nums1[i] = nums2[right]
			right--
		}
	}
}

// 直接塞进去，然后进行排序
func merge2(nums1 []int, m int, nums2 []int, n int) {
	for _, i2 := range nums2 {
		nums1[m] = i2
		m++
	}
	sort.Ints(nums1)
}
