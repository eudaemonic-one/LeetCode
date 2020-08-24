var (
    dx = []int{-1, 0, 1, 0}
    dy = []int{0, -1, 0, 1}
)

func orangesRotting(grid [][]int) int {
    m := len(grid)
    if m == 0 {
        return -1
    }
    n := len(grid[0])
    if n == 0 {
        return -1
    }
    
    freshOranges := 0
    queue := make([][]int, 0)
    
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] == 2 {
                queue = append(queue, []int{r, c})
            } else if grid[r][c] == 1 {
                freshOranges++
            }
        }
    }
    
    queue = append(queue, []int{-1, -1})
    minutesElapsed := -1
    
    for len(queue) > 0 {
        pair := queue[0]
        queue = queue[1:]
        r, c := pair[0], pair[1]
        if r == -1 && c == -1 {
            minutesElapsed++
            if len(queue) > 0 {
                queue = append(queue, []int{-1, -1})
            }
        } else {
            for i := 0; i < 4; i++ {
                nr := r + dx[i]
                nc := c + dy[i]
                if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
                    grid[nr][nc] = 2
                    freshOranges--
                    queue = append(queue, []int{nr, nc})
                }
            }
        }
    }
        
    if freshOranges > 0 {
        return -1
    }
    return minutesElapsed
}
