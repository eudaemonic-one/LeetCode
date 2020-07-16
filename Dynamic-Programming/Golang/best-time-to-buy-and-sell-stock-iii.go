func maxProfit(prices []int) int {
    // dp[i][k][0or1]
    // i: day i
    // k: 0 <= k < K, K is the maximum transaction number
    // 0 means resting and 1 means having stock on hands
    // dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
    // dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
    dp10, dp11 := 0, math.MinInt64
    dp20, dp21 := 0, math.MinInt64
    
    for _, price := range prices {
        dp20 = max(dp20, dp21+price)
        dp21 = max(dp21, dp10-price)
        dp10 = max(dp10, dp11+price)
        dp11 = max(dp11, -price)
    }
    
    return dp20
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
