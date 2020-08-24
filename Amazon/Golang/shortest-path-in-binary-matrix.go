type Cell struct {
    r int
    c int
    depth int
}

func shortestPathBinaryMatrix(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 || grid[0][0] == 1 {
        return -1
    }
    
    m := len(grid)
    n := len(grid[0])
    dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
    dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}
    
    queue := make([]Cell, 0)
    queue = append(queue, Cell{0, 0, 1})
    
    for len(queue) > 0 {
        cell := queue[0]
        queue = queue[1:]
        r := cell.r
        c := cell.c
        depth := cell.depth
        if r == m-1 && c == n-1 {
            return depth
        }
        
        for i := 0; i < len(dx); i++ {
            nr := r + dx[i]
            nc := c + dy[i]
            if nr < 0 || nr >= m || nc < 0 || nc >= n || grid[nr][nc] != 0 {
                continue
            }
            grid[nr][nc] = 1
            queue = append(queue, Cell{nr, nc, depth+1})
        }
    }
    
    return -1
}
