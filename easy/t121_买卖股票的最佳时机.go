package easy

func maxProfit(prices []int) int {

	// 价格天数可能只有一天
	if len(prices) == 1 {
		return 0
	}
	// 暴力法，，遍历
	var max int
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if (prices[j] - prices[i]) > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}

// 题目的本质是，寻找n的位置的n+m 的位置的最大值，求差值
func maxProfit2(prices []int) int {
	var minPrices = prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrices {
			minPrices = prices[i]
		}
		if (prices[i] - minPrices) > max {
			max = prices[i] - minPrices
		}
	}
	return max
}
