package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit_ := 0

	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit_ {
			maxProfit_ = prices[i] - minPrice
		}
	}
	return maxProfit_
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5, 6}))
}
