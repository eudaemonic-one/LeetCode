func maxProfit(prices []int) int {
    res := 0
    minPrice := math.MaxInt64
    for _, price := range prices {
        if price < minPrice {
            minPrice = price
        } else {
            res += price - minPrice
            minPrice = price
        }
    }
    return res
}
