func solveNQueens(n int) [][]string {
    res := make([][]string, 0)
    cols := make([]bool, n)
    upLeftToDownRight := make([]bool, n + n)
    upRightToDownLeft := make([]bool, n + n)
    board := make([][]byte, n)
    for i := 0; i < n; i++ {
        board[i] = make([]byte, n)
        for j := 0; j < n; j++ {
            board[i][j] = '.'
        }
    }
    backtrack(board, cols, upLeftToDownRight, upRightToDownLeft, n, 0, &res)
    return res
}

func backtrack(board [][]byte, cols, upLeftToDownRight, upRightToDownLeft []bool,
               n, r int, res *[][]string) {
    if r == n {
        ans := make([]string, n)
        for i := 0; i < n; i++ {
            ans[i] = string(board[i])
        }
        *res = append(*res, ans)
        return
    }
    for c := 0; c < n; c++ {
        if !cols[c] && !upLeftToDownRight[r+c] && !upRightToDownLeft[(r-c+n)%(2*n-1)] {
            board[r][c] = 'Q'
            cols[c] = true
            upLeftToDownRight[r+c] = true
            upRightToDownLeft[(r-c+n)%(2*n-1)] = true
            backtrack(board, cols, upLeftToDownRight, upRightToDownLeft, n, r+1, res)
            board[r][c] = '.'
            cols[c] = false
            upLeftToDownRight[r+c] = false
            upRightToDownLeft[(r-c+n)%(2*n-1)] = false
        }
    }
}
