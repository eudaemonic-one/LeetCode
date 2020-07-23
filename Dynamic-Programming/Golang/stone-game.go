type pair struct {
    fir int
    sec int
}

func stoneGame(piles []int) bool {
    n := len(piles)
    dp := make([][]pair, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]pair, n)
        dp[i][i] = pair{piles[i], 0}
    }
    
    for l := 2; l <= n; l++ {
        for i := 0; i <= n-l; i++ {
            j := l + i - 1
            left := piles[i] + dp[i+1][j].sec
            right := piles[j] + dp[i][j-1].sec
            if left > right {
                dp[i][j].fir = left
                dp[i][j].sec = dp[i+1][j].fir
            } else {
                dp[i][j].fir = right
                dp[i][j].sec = dp[i][j-1].fir
            }
        }
    }
    
    return dp[0][n-1].fir >= dp[0][n-1].sec
}
