package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	// Raveshe kheyli bad moshkele: o(n2)
	// maxProfit := 0
	// for buy := 0; buy < len(prices); buy++ {
	// 	for sell := buy + 1; sell < len(prices); sell++ {
	// 		this_profit := prices[sell] - prices[buy]

	// 		if this_profit > maxProfit {
	// 			maxProfit = this_profit
	// 		}
	// 	}
	// }

	// Raveshe khob
	maxProfit := 0
	buyPrice := prices[0]
	for day := 1; day < len(prices); day++ {
		if buyPrice > prices[day] {
			buyPrice = prices[day]
		} else if maxProfit < prices[day]-buyPrice {
			maxProfit = prices[day] - buyPrice
		}
	}

	fmt.Println(maxProfit)
}
