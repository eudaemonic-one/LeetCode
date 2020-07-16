func maxProfit(k int, prices []int) int {
    n := len(prices)
    if k > n/2 {
        res := 0
        for i := 1; i < n; i++ {
            if diff := prices[i]-prices[i-1]; diff > 0 {
                res += diff
            }
        }
        return res
    }
    
    // dp[i][k][0 or 1]: until day i, the profit after having used k transactions
    // 0 or 1 means whether having stock on hands
    // dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
    // dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
    dp := make([][][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([][]int, k+1)
        for j := 0; j <= k; j++ {
            dp[i][j] = make([]int, 2)
        }
    }
    
    for i := 0; i < n; i++ {
        price := prices[i]
        for j := k; j >= 1; j-- {
            if i == 0 {
                dp[i][j][0] = 0
                dp[i][j][1] = -price
            } else {
                dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+price)
                dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-price)
            }
        }
    }
    
    return dp[n-1][k][0]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
