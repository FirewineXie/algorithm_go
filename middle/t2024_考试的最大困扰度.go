package middle

/*
解题思路：
首先看题，题目一定存在最长T  或者 最长F
那么 先区分两个计算，
然而计算方式是一样的，

难点：
滑动窗口的流动
长度的计算，

*/

// label :  字符串 ，二分查找，前缀和 ，滑动窗口
func maxConsecutiveAnswers(answerKey string, k int) int {
	return maxLength(maxConsecutiveChar(answerKey, k, 'T'), maxConsecutiveChar(answerKey, k, 'F'))

}

// 滑动窗口，类似与 队列
func maxConsecutiveChar(answerKey string, k int, ch byte) (ans int) {
	left, sum := 0, 0
	for right := range answerKey {
		if answerKey[right] != ch {
			sum++
		}
		for sum > k {
			if answerKey[left] != ch {
				sum--
			}
			left++
		}
		ans = maxLength(ans, right-left+1)
	}
	return
}

func maxLength(a, b int) int {
	if b > a {
		return b
	}
	return a
}
