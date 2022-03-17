package sorted

import "fmt"

// 快排，递归排序，然后合并
func quickSort(arr []int) {

	qsort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func qsort(arr []int, low, high int) {
	if low >= high {
		return
	}

	pivot := partition(arr, low, high) // 将数组分为两部分
	qsort(arr, low, pivot-1)
	qsort(arr, pivot+1, high)

}
func partition(arr []int, low, high int) int {
	pivot := arr[low] //基准
	for low < high {
		for low < high && arr[high] >= pivot {
			high -= 1
		}
		arr[low] = arr[high]
		for (low < high) && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low]
	}
	// 扫描完成，基准到位
	arr[low] = pivot

	return low
}
