func totalNQueens(n int) int {
    cols := make([]bool, n)
    upLeftToDownRight := make([]bool, n + n)
    upRightToDownLeft := make([]bool, n + n)
    res := 0
    backtrack(cols, upLeftToDownRight, upRightToDownLeft, n, 0, &res)
    return res
}

func backtrack(cols, upLeftToDownRight, upRightToDownLeft []bool, n, r int, res *int) {
    if r == n {
        *res += 1
        return
    }
    for c := 0; c < n; c++ {
        if !cols[c] && !upLeftToDownRight[r+c] && !upRightToDownLeft[(r-c+n)%(2*n-1)] {
            cols[c] = true
            upLeftToDownRight[r+c] = true
            upRightToDownLeft[(r-c+n)%(2*n-1)] = true
            backtrack(cols, upLeftToDownRight, upRightToDownLeft, n, r+1, res)
            cols[c] = false
            upLeftToDownRight[r+c] = false
            upRightToDownLeft[(r-c+n)%(2*n-1)] = false
        }
    }
}
