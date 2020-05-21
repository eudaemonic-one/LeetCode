func solveSudoku(board [][]byte)  {
    if board == nil || len(board) != 9 || len(board) != 9 {
        return
    }
    
    // Initialize state space
    rows := make([][]bool, 9)
    cols := make([][]bool, 9)
    subboxes := make([][]bool, 9)
    for i := 0; i < 9; i++ {
        rows[i] = make([]bool, 9)
        cols[i] = make([]bool, 9)
        subboxes[i] = make([]bool, 9)
    }
    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if digit := board[r][c]; digit != '.' {
                digit -= '1'
                rows[r][digit] = true
                cols[c][digit] = true
                subboxes[r/3*3+c/3][digit] = true
            }
        }
    }
    
    backtracking(board, rows, cols, subboxes, 0, 0)
}

func backtracking(board [][]byte, rows, cols, subboxes [][]bool, r, c int) bool {
    // Never alter occupied tiles
    for r < 9 && board[r][c] != '.' {
        c++
        if c == 9 {
            r += 1
            c = 0
        }
    }
    
    // Solution found
    if r == 9 {
        return true
    }
        
    // Search for valid digit that fits current tile
    for digit := 1; digit <= 9; digit++ {
        if !isPlaceble(rows, cols, subboxes, r, c, digit) {
            continue
        }
        
        placeDigit(board, rows, cols, subboxes, r, c, digit)
        
        if backtracking(board, rows, cols, subboxes, r, c) {
            return true
        }
        
        removeDigit(board, rows, cols, subboxes, r, c, digit)
    }
    
    return false
}

func isPlaceble(rows, cols, subboxes [][]bool, r, c, digit int) bool {
    digit -= 1
    return !rows[r][digit] && !cols[c][digit] && !subboxes[r/3*3+c/3][digit]
}

func placeDigit(board [][]byte, rows, cols, subboxes [][]bool, r, c, digit int) {
    digit -= 1
    board[r][c] = byte(digit + '1')
    rows[r][digit] = true
    cols[c][digit] = true
    subboxes[r/3*3+c/3][digit] = true
}

func removeDigit(board [][]byte, rows, cols, subboxes [][]bool, r, c, digit int) {
    digit -= 1
    board[r][c] = '.'
    rows[r][digit] = false
    cols[c][digit] = false
    subboxes[r/3*3+c/3][digit] = false
}
