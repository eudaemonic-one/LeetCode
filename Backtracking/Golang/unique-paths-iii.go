func uniquePathsIII(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    res := 0
    sx, sy, ex, ey := 0, 0, 0, 0
    spaces := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 {
                sx, sy = i, j
            } else if grid[i][j] == 2 {
                ex, ey = i, j
            }
            if grid[i][j] != -1 {
                spaces += 1
            }
        }
    }
    dfs(grid, sx, sy, ex, ey, spaces, &res)
    return res
}

func dfs(grid [][]int, x, y, ex, ey, spaces int, res *int) {
    spaces--
    
    // out of grid or hit obstacles
    if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == -1 {
        return
    }
    
    if spaces <= 0 {
        // reach ending square
        if x == ex && y == ey {
            *res += 1
        }
        return
    }
    
    grid[x][y] = -1
    defer func() {
       grid[x][y] = 0 
    }()
    
    dfs(grid, x, y-1, ex, ey, spaces, res)
    dfs(grid, x, y+1, ex, ey, spaces, res)
    dfs(grid, x-1, y, ex, ey, spaces, res)
    dfs(grid, x+1, y, ex, ey, spaces, res)
}
