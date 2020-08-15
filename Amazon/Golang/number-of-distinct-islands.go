func numDistinctIslands(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    shapeSet := make(map[string]bool)
    visited := make([][]bool, len(grid))
    for i := 0; i < len(grid); i++ {
        visited[i] = make([]bool, len(grid[0]))
    }
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[0]); c++ {
            shape := make([]byte, 0)
            dfs(&shape, grid, visited, r, c, '0')
            if len(shape) > 0 {
                shapeSet[string(shape)] = true
            }
        }
    }
    return len(shapeSet)
}

func dfs(path *[]byte, grid [][]int, visited [][]bool, r, c int, direction byte) {
    if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || visited[r][c] == true || grid[r][c] != 1 {
        return
    }
    visited[r][c] = true
    *path = append(*path, direction)
    dfs(path, grid, visited, r+1, c, '1')
    dfs(path, grid, visited, r-1, c, '2')
    dfs(path, grid, visited, r, c+1, '3')
    dfs(path, grid, visited, r, c-1, '4')
    *path = append(*path, '0')
}
