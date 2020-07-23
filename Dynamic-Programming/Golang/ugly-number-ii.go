func nthUglyNumber(n int) int {
    if n <= 5 {
        return n
    }
    
    dp := []int{1}
    twos, threes, fives := 0, 0, 0
    for i := 1; i < n; i++ {
        ugly := min(dp[twos] * 2, min(dp[threes] * 3, dp[fives] * 5))
        dp = append(dp, ugly)
        
        if ugly == dp[twos] * 2 {
            twos++
        }
        if ugly == dp[threes] * 3 {
            threes++
        }
        if ugly == dp[fives] * 5 {
            fives++
        }
    }
    return dp[n-1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
