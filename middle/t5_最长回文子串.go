package middle

import "strings"

// 动态规划
func longestPalindrome(s string) string {

	length := len(s)
	if length < 2 {
		return s
	}

	maxLen := 1
	begin := 0

	// dp[i][j] 表示 s[i,j]是否是回文串
	var dp [][]bool
	// 初始化 ： 所有长度为1 的子串都是回文串
	for i := 0; i < length; i++ {
		var tmp []bool
		for j := 0; j < length; j++ {
			tmp = append(tmp, true)
		}
		dp = append(dp, tmp)
	}

	split := strings.Split(s, "")

	// 递推开始
	for l := 2; l < length; l++ {
		// 枚举左边界，左边界的上限设置可以宽松一些
		for i := 0; i < length; i++ {
			// 有 l 和 i 可以确定右边界， 即 j- i + 1 = l
			j := l + i - 1
			// 右边界越界，退出当前循环
			if j >= length {
				break
			}
			if split[i] != split[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			// 只要 dp[i][L] == true 成立，就表示子串 s[i..L] 是回文，此时记录回文长度和起始位
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	if begin+maxLen == len(split)-1 {
		return strings.Join(split[begin:begin+maxLen], "")
	}
	return strings.Join(split[begin:begin+maxLen+1], "")

}
