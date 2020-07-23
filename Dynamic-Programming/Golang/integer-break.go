func integerBreak(n int) int {
    // dp[i]: the maximum factor product until i
    dp := make([]int, n+1)
    dp[1] = 1
    
    for i := 2; i <= n; i++ {
        for j := 1; j <= i/2; j++ {
            dp[i] = max(dp[i], max(j, dp[j]) * max(i-j, dp[i-j]))
        }
    }
    
    return dp[n]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
