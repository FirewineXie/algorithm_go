package easy

import "sort"

// 先排序，然后进行双指针扫
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	nums1S, nums2S := 0, 0
	var result []int
	for true {
		if nums1[nums1S] == nums2[nums2S] {
			result = append(result, nums1[nums1S])
			nums1S++
			nums2S++
		} else if nums1[nums1S] > nums2[nums2S] {
			nums2S++
		} else {
			nums1S++
		}
		if nums1S == len(nums1) || nums2S == len(nums2) {
			break
		}
	}
	return result
}

//进阶：
//
//    如果给定的数组已经排好序呢？你将如何优化你的算法？
//    如果 nums1 的大小比 nums2 小，哪种方法更优？
//    如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
//

func intersect2(nums1 []int, nums2 []int) []int {
	hashTable := map[int]int{}
	var result []int
	for _, v := range nums1 {
		if _, ok := hashTable[v]; ok {
			hashTable[v] += 1
		} else {
			hashTable[v] = 1
		}
	}

	for _, v := range nums2 {
		if num, ok := hashTable[v]; ok && num > 0 {
			hashTable[v]--
			result = append(result, v)
		}
	}
	return result
}
