package middle

// 解题思路： 遍历两个数组，如果如果一个数组还有剩余，那么直接拼接到后面
// 要对于边界条件，进行判断，各种条件进行多方面考虑
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{}
	tt := &head
	temp := 0
	for true {
		val := temp
		temp = 0
		if l1 != nil {
			val += l1.Val
			if l1.Next != nil {
				l1 = l1.Next
			} else {
				l1 = nil
			}
		}
		if l2 != nil {
			val += l2.Val
			if l2.Next != nil {
				l2 = l2.Next
			} else {
				l2 = nil
			}
		}

		d := val % 10
		if val >= 10 {
			temp = val / 10
		}
		tt.Val = d
		if l1 != nil || l2 != nil || temp != 0 {
			tt.Next = &ListNode{}
			tt = tt.Next
		}

		if l1 == nil && l2 == nil && temp == 0 {
			break
		}
	}

	return &head
}
