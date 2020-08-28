func countBattleships(board [][]byte) int {
    m := len(board)
    if m == 0 {
        return 0
    }
    n := len(board[0])
    if n == 0 {
        return 0
    }
    
    count := 0
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == '.' {
                continue
            }
            if j > 0 && board[i][j-1] == 'X' {
                continue
            }
            if i > 0 && board[i-1][j] == 'X' {
                continue
            }
            count++
        }
    }
    
    return count
}
