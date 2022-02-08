package easy

import "sort"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// 1. 初步断定是二分查找，递归寻找就可以了

func firstBadVersion(n int) int {

	low, high, mid := 1, n, 0

	for low < high {
		if !isBadVersion(low) && isBadVersion(high) {
			// 前 false , 后 true ，证明firstVersion在这个区间，二分缩短区间查找
			mid = (low + high) / 2
			if isBadVersion(mid) {
				// 中间为true ，往前查找
				high = mid
			} else {
				// 中间为false ，往后查找
				low = mid + 1
			}
		} else {
			// 如果两个都是true
			return low
		}
	}

	return low
}

func isBadVersion(n int) bool {
	return true
}

func firstBadVersion2(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}
