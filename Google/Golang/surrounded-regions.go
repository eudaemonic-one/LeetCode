func solve(board [][]byte) {
    if len(board) == 0 || len(board[0]) == 0 {
        return
    }
    m, n := len(board), len(board[0])
    for i := 0; i < m; i++ {
        helper(board, i, 0, m, n)
        helper(board, i, n-1, m, n)
    }
    for j := 0; j < n; j++ {
        helper(board, 0, j, m, n)
        helper(board, m-1, j, m, n)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            } else if board[i][j] == '#' {
                board[i][j] = 'O'
            }
        }
    }
}

func helper(board [][]byte, x, y, m, n int) {
    if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
        return
    }
    board[x][y] = '#'
    helper(board, x-1, y, m, n)
    helper(board, x, y-1, m, n)
    helper(board, x, y+1, m, n)
    helper(board, x+1, y, m, n)
}
