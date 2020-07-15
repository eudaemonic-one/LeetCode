func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    if m * n == 0 {
        return m + n
    }
    
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
    }
    for j := 0; j < n+1; j++ {
        dp[0][j] = j
    }
    for i := 0; i < m+1; i++ {
        dp[i][0] = i
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            left := dp[i][j+1] + 1
            down := dp[i+1][j] + 1
            leftDown := dp[i][j]
            if word1[i] != word2[j] {
                leftDown += 1
            }
            dp[i+1][j+1] = min(left, min(down, leftDown))
        }
    }
    
    return dp[m][n]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
