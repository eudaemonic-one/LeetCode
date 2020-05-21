func getMaximumGold(grid [][]int) int {
    res := 0
    m, n := len(grid), len(grid[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] > 0 {
                dfs(grid, m, n, i, j, 0, &res)
            }
        }
    }
    return res
}

func dfs(grid [][]int, m, n, i, j, sum int, res *int) {
    if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
        return
    }
    
    gold := grid[i][j]
    sum += gold
    grid[i][j] = 0
    
    dfs(grid, m, n, i, j-1, sum, res)
    dfs(grid, m, n, i, j+1, sum, res)
    dfs(grid, m, n, i-1, j, sum, res)
    dfs(grid, m, n, i+1, j, sum, res)
    
    grid[i][j] = gold
    
    *res = max(*res, sum)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
