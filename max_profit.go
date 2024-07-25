package main

import "fmt"

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    minPrice := prices[0]
    maxProfit := 0

    for _, price := range prices[1:] {
        if price < minPrice {
            minPrice = price
        } else if price-minPrice > maxProfit {
            maxProfit = price - minPrice
        }
    }

    return maxProfit
}

func main() {
    testCases := [][]int{
        {10, 9, 6, 5, 15},
        {7, 8, 3, 10, 8},
        {5, 12, 11, 12, 10},
        {7, 18, 27, 10, 29},
        {20, 17, 15, 14, 10},
    }

    for _, prices := range testCases {
        fmt.Println(maxProfit(prices))
    }
}
