const N int = 3

func numberOfPatterns(m int, n int) int {
    res := 0
    used := make([][]bool, N)
    for i := 0; i < N; i++ {
        used[i] = make([]bool, N)
    }
    
    for path := m; path <= n; path++ {
        for i := 0; i < N; i++ {
            for j := 0; j < N; j++ {
                res += backtrack(used, i, j, path-1)
            }
        }
    }
    return res
}

func backtrack(used [][]bool, x, y, path int) int {
    if path == 0 {
        return 1
    }
    sum := 0
    for nx := 0; nx < N; nx++ {
        for ny := 0; ny < N; ny++ {
            if isValidNextPos(used, x, y, nx, ny) {
                used[x][y] = true
                sum += backtrack(used, nx, ny, path-1)
                used[x][y] = false
            }
        }
    }
    return sum
}

func isValidNextPos(used [][]bool, sx, sy, ex, ey int) bool {
    if used[ex][ey] == true {
        return false
    }
    
    dx, dy := abs(sx - ex), abs(sy - ey)
    
    if dx == 0 && dy == 0 {
        return false
    }
        
    if dx == 2 && dy == 0 {
        return used[1][sy]
    }
    
    if dx == 0 && dy == 2 {
        return used[sx][1]
    }
    
    if dx == 2 && dy == 2 {
        return used[1][1]
    }
    
    return true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
