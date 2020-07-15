// Approach #2: DP with Tricks (Compression)

func uniquePaths(m int, n int) int {
    dp := make([]int, n)
    for i := 0; i < n; i++ {
        dp[i] = 1
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[j] = dp[j-1] + dp[j]
        }
    }
    
    return dp[n-1]
}

// Approach #1: DP with Tricks (Parity)

// func uniquePaths(m int, n int) int {
//     dp := make([][]int, 2)
//     for i := 0; i < 2; i++ {
//         dp[i] = make([]int, n)
//         for j := 0; j < n; j++ {
//             dp[i][j] = 1
//         }
//     }
    
//     for i := 1; i < m; i++ {
//         for j := 1; j < n; j++ {
//             dp[i%2][j] = dp[(i-1)%2][j] + dp[i%2][j-1]
//         }
//     }
    
//     return dp[(m-1)%2][n-1]
// }
