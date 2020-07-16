// Approach #2: Space Optimized DP

func numDistinct(s string, t string) int {
    if len(s) < len(t) {
        return 0
    }
    
    m, n := len(s), len(t)
    dp := make([]int, n+1)
    
    for i := m-1; i >= 0; i-- {
        prev := 1
        for j := n-1; j >= 0; j-- {
            old := dp[j]
            if s[i] == t[j] {
                dp[j] += prev
            }
            prev = old
        }
    }
    
    return dp[0]
}

// Approach #1: DP

// func numDistinct(s string, t string) int {
//     if len(s) < len(t) {
//         return 0
//     }
    
//     m, n := len(s), len(t)
//     dp := make([][]int, m+1)
//     for i := 0; i < m+1; i++ {
//         dp[i] = make([]int, n+1)
//         dp[i][n] = 1
//     }
//     for j := 0; j < n; j++ {
//         dp[m][j] = 0
//     }
    
//     for i := m-1; i >= 0; i-- {
//         for j := n-1; j >= 0; j-- {
//             dp[i][j] = dp[i+1][j]
//             if s[i] == t[j] {
//                 dp[i][j] += dp[i+1][j+1]
//             }
//         }
//     }
    
//     return dp[0][0]
// }
