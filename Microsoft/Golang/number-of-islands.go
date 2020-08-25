var count int

func numIslands(grid [][]byte) int {
    count = 0
    
    m := len(grid)
    if m == 0 {
        return 0
    }
    n := len(grid[0])
    if n == 0 {
        return 0
    }
    
    parent := make([]int, m*n)
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] == '1' {
                parent[r*n+c] = r*n + c
                count++
            }
        }
    }
    
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] == '1' {
                grid[r][c] = '0'
                if r-1 >= 0 && grid[r-1][c] == '1' {
                    union(parent, r*n+c, (r-1)*n+c)
                }
                if r+1 < m && grid[r+1][c] == '1' {
                    union(parent, r*n+c, (r+1)*n+c)
                }
                if c-1 >= 0 && grid[r][c-1] == '1' {
                    union(parent, r*n+c, r*n+c-1)
                }
                if c+1 < n && grid[r][c+1] == '1' {
                    union(parent, r*n+c, r*n+c+1)
                }
            }
        }
    }
    
    return count
}

func find(parent []int, x int) int {
    for x != parent[x] {
        x = parent[x]
    }
    return x
}

func union(parent []int, x, y int) {
    rootX := find(parent, x)
    rootY := find(parent, y)
    if rootX != rootY {
        parent[rootX] = rootY
        count--
    }
}
