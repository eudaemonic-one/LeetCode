func maximalSquare(matrix [][]byte) int {
    res := 0
    m := len(matrix)
    if m == 0 {
        return 0
    }
    n := len(matrix[0])
    if n == 0 {
        return 0
    }
    
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if matrix[i-1][j-1] == '1' {
                dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
                if dp[i][j] > res {
                    res = dp[i][j]
                }
            }
            
        }
    }
        
    return res * res
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
