func maxProfit(prices []int) int {
    if len(prices) <= 1 {
        return 0
    }
    
    // dp[i][j][k]: until day i, the maximum profit when on hand stock is j and k is whether we buy or sell right now
    // we see i only use i-1 and thus the i-dimension can be ignored
    // dp[1][0] = max(dp[1][0], dp[0][1])
    // dp[0][0] = max(dp[0][0], dp[1][1])
    // dp[0][1] = dp[1][0] - prices[i]
    // dp[1][1] = dp[1][0] + prices[i]
    have0buy0 := 0
    have0buy1 := -prices[0]
    have1sell0 := -prices[0]
    have1sell1 := 0
    
    for i := 1; i < len(prices); i++ {
        have1sell0 = max(have1sell0, have0buy1)
        have0buy1 = have0buy0 - prices[i]
        have0buy0 = max(have0buy0, have1sell1)
        have1sell1 = have1sell0 + prices[i]
    }
    
    return max(have0buy0, have1sell1)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
