func minCost(costs [][]int) int {
    if len(costs) == 0 {
        return 0
    }
    
    dp := make([]int, 3)
    
    for i := 0; i < len(costs); i++ {
        row := make([]int, 3)
        copy(row, costs[i])
        row[0] += min(dp[1], dp[2])
        row[1] += min(dp[0], dp[2])
        row[2] += min(dp[0], dp[1])
        copy(dp, row)
    }
    
    return min(dp[0], min(dp[1], dp[2]))
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
