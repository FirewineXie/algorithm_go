package easy

func reverseString(s []byte) {

	left, right := 0, len(s)-1
	var temp byte
	for true {

		if right < left || right == left {
			break
		}
		temp = s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}
