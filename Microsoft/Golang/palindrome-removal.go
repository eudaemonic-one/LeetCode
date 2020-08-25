func minimumMoves(arr []int) int {
    n := len(arr)
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = n
        }
        dp[i][i] = 1
        if i < n-1 {
            if arr[i] == arr[i+1] {
                dp[i][i+1] = 1
            } else {
                dp[i][i+1] = 2
            }
        }
    }
    
    for subLen := 3; subLen <= n; subLen++ {
        for i, j := 0, subLen-1; j < n; i, j = i+1, j+1 {
            if arr[i] == arr[j] {
                dp[i][j] = dp[i+1][j-1]
            }
            for k := i; k < j; k++ {
                dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
            }
        }
    }
    
    return dp[0][n-1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
