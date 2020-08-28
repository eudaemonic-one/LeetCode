const sentinel = byte('0')

func exist(board [][]byte, word string) bool {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if dfs(board, word, i, j, 0) {
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, word string, i, j, idx int) bool {
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || word[idx] != board[i][j] {
        return false
    }
    
    if idx == len(word)-1 {
        return true
    }
    
    ch := board[i][j]
    board[i][j] = sentinel
    defer func() {
        board[i][j] = ch
    }()
    
    if dfs(board, word, i-1, j, idx+1) { // move up
        return true
    }
    if dfs(board, word, i+1, j, idx+1) { // move down
        return true
    }
    if dfs(board, word, i, j-1, idx+1) { // move left
        return true
    }
    if dfs(board, word, i, j+1, idx+1) { // move right
        return true
    }
    
    return false
}
