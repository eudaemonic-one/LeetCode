func exist(board [][]byte, word string) bool {
    if len(board) == 0 || len(word) == 0 {
        return false
    }
    visited := make([][]bool, len(board))
    for i := 0; i < len(board); i++ {
        visited[i] = make([]bool, len(board[i]))
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if dfs(board, word, &visited, i, j, 0) {
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, word string, visited *[][]bool, i, j, idx int) bool {
    if word[idx] != board[i][j] {
        return false
    }
    
    if idx + 1 == len(word) {
        return true
    }
    
    (*visited)[i][j] = true
    defer func() {
        (*visited)[i][j] = false
    }()
    
    // Move left
    if j > 0 && !(*visited)[i][j-1] {
        if dfs(board, word, visited, i, j-1, idx+1) {
            return true
        }
    }
    // Move right
    if j + 1 < len(board[i]) && !(*visited)[i][j+1] {
        if dfs(board, word, visited, i, j+1, idx+1) {
            return true
        }
    }
    // Move up
    if i > 0 && !(*visited)[i-1][j] {
        if dfs(board, word, visited, i-1, j, idx+1) {
            return true
        }
    }
    // Move down
    if i + 1 < len(board) && !(*visited)[i+1][j] {
        if dfs(board, word, visited, i+1, j, idx+1) {
            return true
        }
    }
    
    return false
}
