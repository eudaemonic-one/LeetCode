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

func dfs(grid [][]int, m, n, i, j, sum int, res *int) bool {
    if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
        return false
    }
    
    gold := grid[i][j]
    sum += gold
    grid[i][j] = 0
    defer func() {
        grid[i][j] = gold
    }()
    
    flag := false
    if dfs(grid, m, n, i, j-1, sum, res) {
        flag = true
    }
    if dfs(grid, m, n, i, j+1, sum, res) {
        flag = true
    }
    if dfs(grid, m, n, i-1, j, sum, res) {
        flag = true
    }
    if dfs(grid, m, n, i+1, j, sum, res) {
        flag = true
    }
    
    if flag == false {
        *res = max(*res, sum)
        return true
    }
    
    return false
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
