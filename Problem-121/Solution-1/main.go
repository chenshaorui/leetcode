package Solution_1

func maxProfit(prices []int) int {
	maximumProfit := 0

	days := len(prices)
	if days > 0 {
		minPrice := prices[0]

		for i := 1; i < days; i++ {
			price := prices[i]

			if price < minPrice {
				minPrice = price
			} else {
				profit := price - minPrice
				if profit > maximumProfit {
					maximumProfit = profit
				}
			}
		}
	}

	return maximumProfit
}
