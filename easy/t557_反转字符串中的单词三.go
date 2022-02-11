package easy

import "strings"

func reverseWords(s string) string {
	if len(s) == 1 {
		return s
	}
	result := new(strings.Builder)
	result.Grow(len(s))

	left, right := 0, 1
	var tempNum int
	for true {
		// 右指针 到达空格处 或者字符串长度，开始反转。并且到达终点，反转完后退出
		if s[right] == ' ' {
			tempNum = right // 记录当前空格位置
			for left != right {
				right--
				result.WriteByte(s[right])

			}
			right = tempNum
			result.WriteByte(' ') // 补充空格
			right++
			left = right
		}
		// 当到达最后
		if right == len(s)-1 {
			for true {
				result.WriteByte(s[right])
				if left == right {
					break
				}
				right--

			}
			right = tempNum
			break
		}

		right++
	}

	return result.String()
}
