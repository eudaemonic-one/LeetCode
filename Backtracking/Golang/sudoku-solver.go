const (
    digits = "123456789"
)

func solveSudoku(board [][]byte)  {
    if board == nil || len(board) != 9 || len(board) != 9 {
        return
    }
    
    // Initialize state space
    rowsDigitSets := make([]map[byte]bool, 9)
    for i := 0; i < 9; i++ {
        rowsDigitSets[i] = make(map[byte]bool, 3)
    }
    columnsDigitSets := make([]map[byte]bool, 9)
    for i := 0; i < 9; i++ {
        columnsDigitSets[i] = make(map[byte]bool, 3)
    }
    subboxesDigitSets := make([][]map[byte]bool, 3)
    for i := 0; i < 3; i++ {
        subboxesDigitSets[i] = make([]map[byte]bool, 3)
        for j := 0; j < 3; j++ {
            subboxesDigitSets[i][j] = make(map[byte]bool)
        }
    }
    res := make([][]byte, 9)
    for i := 0; i < 9; i++ {
        res[i] = make([]byte, 9)
    }
    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if board[r][c] != '.' {
                digit := board[r][c]
                rowsDigitSets[r][digit] = true
                columnsDigitSets[c][digit] = true
                subboxesDigitSets[r/3][c/3][digit] = true
            }
        }
    }
    backtracking(board, rowsDigitSets, columnsDigitSets, subboxesDigitSets, 0, 0, res)
    
    // Copy res to board in-place
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            board[i][j] = res[i][j]
        }
    }
}

func backtracking(board [][]byte, rowsDigitSets, columnsDigitSets []map[byte]bool,
                  subboxesDigitSets [][]map[byte]bool, r, c int, res [][]byte) {
    // Solution found
    if r == 9 {
        for i := 0; i < 9; i++ {
            for j := 0; j < 9; j++ {
                res[i][j] = board[i][j]
            }
        }
        return
    }
    
    // Never alter occupied tiles
    if board[r][c] != '.' {
        nr, nc := nextPos(r, c)
        backtracking(board, rowsDigitSets, columnsDigitSets, subboxesDigitSets, nr, nc, res)
        return
    }
    
    // Search for valid digit that fits current tile and shift to next position sequentially
    for k := 0; k < len(digits); k++ {
        digit := digits[k]
        
        if !isPlaceble(rowsDigitSets, columnsDigitSets, subboxesDigitSets, r, c, digit) {
            continue
        }
        
        placeDigit(board, rowsDigitSets, columnsDigitSets, subboxesDigitSets, r, c, digit)
        
        nr, nc := nextPos(r, c)
        
        backtracking(board, rowsDigitSets, columnsDigitSets, subboxesDigitSets, nr, nc, res)
        
        removeDigit(board, rowsDigitSets, columnsDigitSets, subboxesDigitSets, r, c, digit)
    }
}

func isPlaceble(rowsDigitSets, columnsDigitSets []map[byte]bool,
                subboxesDigitSets [][]map[byte]bool,
                r, c int, digit byte) bool {
    if flag, ok := rowsDigitSets[r][digit]; ok && flag {
        return false
    } else if flag, ok := columnsDigitSets[c][digit]; ok && flag {
        return false
    } else if flag, ok := subboxesDigitSets[r/3][c/3][digit]; ok && flag {
        return false
    }
    return true
}

func placeDigit(board [][]byte, rowsDigitSets, columnsDigitSets []map[byte]bool,
                  subboxesDigitSets [][]map[byte]bool, r, c int, digit byte) {
    board[r][c] = digit
    rowsDigitSets[r][digit] = true
    columnsDigitSets[c][digit] = true
    subboxesDigitSets[r/3][c/3][digit] = true
}

func removeDigit(board [][]byte, rowsDigitSets, columnsDigitSets []map[byte]bool,
                  subboxesDigitSets [][]map[byte]bool, r, c int, digit byte) {
    board[r][c] = '.'
    rowsDigitSets[r][digit] = false
    columnsDigitSets[c][digit] = false
    subboxesDigitSets[r/3][c/3][digit] = false
}

func nextPos(r, c int) (int, int) {
    nr, nc := r, c + 1
    if nc == 9 {
        nr, nc = r + 1, 0
    }
    return nr, nc
}
