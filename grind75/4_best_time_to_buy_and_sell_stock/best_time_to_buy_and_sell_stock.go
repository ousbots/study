package best_time_to_buy_and_sell_stock

// Find the maximum profit that can be made buying and selling at the given prices.
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0
	lowest := prices[0]

	for _, sell := range prices {
		if sell < lowest {
			lowest = sell
		}

		if sell-lowest > maxProfit {
			maxProfit = sell - lowest
		}
	}

	return maxProfit
}

func maxProfitBruteForce(prices []int) int {
	maxProfit := 0

	for i, buy := range prices {
		for _, sell := range prices[i+1:] {
			if sell-buy > maxProfit {
				maxProfit = sell - buy
			}
		}
	}

	return maxProfit
}
