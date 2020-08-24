type TicTacToe struct {
    n            int
    rows         []int
    cols         []int
    diagonal     int
    antiDiagonal int
    winner       int
}


/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
    return TicTacToe{n, make([]int, n), make([]int, n), 0, 0, 0}
}


/** Player {player} makes a move at ({row}, {col}).
        @param row The row of the board.
        @param col The column of the board.
        @param player The player, can be either 1 or 2.
        @return The current winning condition, can be either:
                0: No one wins.
                1: Player 1 wins.
                2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
    if this.winner != 0 {
        return this.winner
    }
    
    var toAdd int
    if player == 1 {
        toAdd = 1
    } else {
        toAdd = -1
    }
    
    this.rows[row] += toAdd
    this.cols[col] += toAdd
    if row == col {
        this.diagonal += toAdd
    }
    if row+col == this.n-1 {
        this.antiDiagonal += toAdd
    }
    
    if abs(this.rows[row]) == this.n || abs(this.cols[col]) == this.n ||
            abs(this.diagonal) == this.n || abs(this.antiDiagonal) == this.n {
        this.winner = player
        return player
    }
    
    return 0
}


func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
