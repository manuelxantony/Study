package main

// 29.26 %
// 2 pointer
func maxProfit(prices []int) int {
    max := 0
    l := 0
    r := 1
	lp := len(prices)
    for r < lp {
        if prices[r] > prices[l]{
			currentProfit := prices[r] - prices[l]
			if  currentProfit > max{
				max = currentProfit
			}
		} else {
			l = r
		}
		r += 1
    }
    return max
}

// much faster / 98.5%
// straight logic
func maxProfit2(prices []int) int {
    maxProfit := 0
    priceMin := prices[0]

    for i:=0; i < len(prices); i++ {
        if priceMin > prices[i]{
            priceMin = prices[i]
        } else{
            currentProfit := prices[i] - priceMin
            if currentProfit > maxProfit {
                maxProfit = currentProfit
            }
        }
    }
    return maxProfit
}