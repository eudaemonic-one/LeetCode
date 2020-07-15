// Approach #2: DP 1D

func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([]int, n)
    
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if i == m-1 && j != n-1 {
                dp[j] = grid[i][j] + dp[j+1]
            } else if i != m-1 && j == n-1 {
                dp[j] = grid[i][j] + dp[j]
            } else if i != m-1 && j != n-1 {
                dp[j] = grid[i][j] + min(dp[j], dp[j+1])
            } else {
                dp[j] =  grid[i][j]
            }
        }
    }
    
    return dp[0]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

// Approach #1: Bottom-up DP 2D

// func minPathSum(grid [][]int) int {
//     m, n := len(grid), len(grid[0])
//     dp := make([][]int, m+1)
//     for i := 0; i < m+1; i++ {
//         dp[i] = make([]int, n+1)
//     }
    
//     for i := m-1; i >= 0; i-- {
//         for j := n-1; j >= 0; j-- {
//             dp[i][j] = grid[i][j]
//             if i == m-1 {
//                 dp[i][j] += dp[i][j+1]
//             } else if j == n-1 {
//                 dp[i][j] += dp[i+1][j]
//             } else {
//                 dp[i][j] += min(dp[i+1][j], dp[i][j+1])
//             }
//         }
//     }
    
//     return dp[0][0]
// }

// func min(x, y int) int {
//     if x < y {
//         return x
//     }
//     return y
// }
