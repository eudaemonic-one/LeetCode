func numSquares(n int) int {
    maxRoot := int(math.Sqrt(float64(n)))
    dp := make([]int, n+1)
    for sum := 1; sum <= n; sum++ {
        dp[sum] = math.MaxInt64
        for root := 1; root <= maxRoot; root++ {
            if sum - root*root < 0 {
                break
            }
            if tmp := dp[sum-root*root] + 1; tmp < dp[sum] {
                dp[sum] = tmp
            }
        }
    }
    return dp[n]
}
