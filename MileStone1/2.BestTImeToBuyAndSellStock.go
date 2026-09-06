package main

func bestTimeToBuyAndSellStock(prices []int) int {
	r := 1
	l := 0
	maxp := 0
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxp = max(maxp, profit)
		} else {
			l = r
		}
		r++
	}
	println("2.", maxp)
	return maxp
}
