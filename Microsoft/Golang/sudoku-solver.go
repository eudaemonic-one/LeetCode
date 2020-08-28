func solveSudoku(board [][]byte)  {
    if board == nil || len(board) != 9 || len(board[0]) != 9 {
        return
    }

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
                num := int(digit-'1')
                rows[r][num] = true
                cols[c][num] = true
                subboxes[r/3*3+c/3][num] = true
            }
        }
    }
    
    dfs(board, 0, 0, rows, cols, subboxes)
}

func dfs(board [][]byte, r, c int, rows, cols, subboxes [][]bool) bool {
    for r < 9 && board[r][c] != '.' {
        c++
        if c == 9 {
            r++
            c = 0
        }
    }
    
    if r == 9 {
        return true // solution found
    }
    
    for digit := 1; digit <= 9; digit++ {
        if !isPlaceble(rows, cols, subboxes, r, c, digit-1) {
            continue
        }
        
        board[r][c] = byte(digit+'0')
        rows[r][digit-1] = true
        cols[c][digit-1] = true
        subboxes[r/3*3+c/3][digit-1] = true
        
        if dfs(board, r, c, rows, cols, subboxes) {
            return true
        }
        
        board[r][c] = '.'
        rows[r][digit-1] = false
        cols[c][digit-1] = false
        subboxes[r/3*3+c/3][digit-1] = false
    }
    
    return false
}

func isPlaceble(rows, cols, subboxes [][]bool, r, c, digit int) bool {
    return !rows[r][digit] && !cols[c][digit] && !subboxes[r/3*3+c/3][digit]
}

