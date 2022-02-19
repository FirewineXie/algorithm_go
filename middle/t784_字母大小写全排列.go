package middle

func letterCasePermutation(s string) []string {
	var ans []string
	backtrace([]byte(s), 0, &ans)
	return ans
}

func backtrace(str []byte, start int, ans *[]string) {
	*ans = append(*ans, string(str))
	for i := start; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			str[i] -= 32
			backtrace(str, i+1, ans)
			str[i] += 32
		}
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] += 32
			backtrace(str, i+1, ans)
			str[i] -= 32
		}
	}
}
