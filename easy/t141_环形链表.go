package easy

// 使用额外的空间计算
func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	data := map[*ListNode]int{}
	var i int
	for head.Next != nil {
		if _, ok := data[head]; ok {
			return true
		}
		i++
		data[head] = i
		head = head.Next
	}

	return false
}

//// 使用O(1) 空间，就是常量，解决问题
//func hasCycleV1(head *ListNode) bool {
//	if head == nil {
//		return false
//	}
//	down := head
//
//	for down.Next != nil {
//
//	}
//
//}
