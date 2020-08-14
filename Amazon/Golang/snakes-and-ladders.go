func snakesAndLadders(board [][]int) int {
    if len(board) == 0 || len(board[0]) == 0 {
        return 0
    }
    
    N := len(board)
    dist := make(map[int]int)
    dist[1] = 0
    queue := []int{1}
    
    for len(queue) > 0 {
        s := queue[0]
        queue = queue[1:]
        if s == N*N {
            return dist[s]
        }
        
        for s2 := s+1; s2 <= min(s+6, N*N); s2++ {
            r, c := getCoordinate(s2, N)
            s2Final := s2
            if board[r][c] != -1 {
                s2Final = board[r][c]
            }
            if _, ok := dist[s2Final]; !ok {
                dist[s2Final] = dist[s] + 1
                queue = append(queue, s2Final)
            }
        }
    }
    
    return -1
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func getCoordinate(s, N int) (int, int) {
    var row, col int
    quot := (s-1) / N
    rem := (s-1) % N    
    row = N - 1 - quot
    if row % 2 != N % 2 {
        col = rem
    } else {
        col = N - 1 - rem
    }
    return row, col
}
