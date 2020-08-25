func validTicTacToe(board []string) bool {
    countX, countO := 0, 0
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == 'X' {
                countX++
            } else if board[i][j] == 'O' {
                countO++
            }
        }
    }
    
    if countX != countO && countX != countO+1 {
        return false
    }
    
    if win(board, 'X') && countX != countO+1 {
        return false
    }
    
    if win(board, 'O') && countX != countO {
        return false
    }
    
    return true
}

func win(board []string, player byte) bool {
    for i := 0; i < 3; i++ {
        if player == board[0][i] && player == board[1][i] && player == board[2][i] {
            return true
        }
        if player == board[i][0] && player == board[i][1] && player == board[i][2] {
            return true
        }
    }
    if player == board[0][0] && player == board[1][1] && player == board[2][2] {
        return true
    }
    if player == board[0][2] && player == board[1][1] && player == board[2][0] {
        return true
    }
    return false
}
