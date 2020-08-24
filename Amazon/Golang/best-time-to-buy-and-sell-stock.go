func maxProfit(prices []int) int {
    res := 0
    minPrice := math.MaxInt64
    for i := 0; i < len(prices); i++ {
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else if diff := prices[i] - minPrice; diff > res {
            res = diff
        }
    }
    return res
}
