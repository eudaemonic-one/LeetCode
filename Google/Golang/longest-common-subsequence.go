func longestCommonSubsequence(text1 string, text2 string) int {
    m := len(text1)
    n := len(text2)
    if m == 0 || n == 0 {
        return 0
    }
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if text1[i] == text2[j] {
                dp[i+1][j+1] = dp[i][j] + 1
            } else {
                dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
            }
        }
    }
    return dp[m][n]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

