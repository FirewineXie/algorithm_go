package easy

// 最简单递归
func addDigits(num int) int {

	if num == 0 {
		return 0
	}

	return dfsAddDigits(num)
}

func dfsAddDigits(num int) int {
	if num < 10 {
		return num
	}
	var sum int
	for true {
		t := num % 10
		sum += t
		if num < 10 {
			break
		}
		num = num / 10
	}
	return dfsAddDigits(sum)
}

//你可以不使用循环或者递归，在 O(1) 时间复杂度内解决这个问题吗？
func addDigitsV1(num int) int {

	return (num-1)%9 + 1
}
