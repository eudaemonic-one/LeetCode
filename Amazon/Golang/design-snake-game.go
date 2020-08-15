type SnakeGame struct {
    height int
    width int
    snake []int
    food [][]int
    foodIndex int
}


/** Initialize your data structure here.
        @param width - screen width
        @param height - screen height 
        @param food - A list of food positions
        E.g food = [[1,1], [1,0]] means the first food is positioned at [1,1], the second is at [1,0]. */
func Constructor(width int, height int, food [][]int) SnakeGame {
    game := SnakeGame{height, width, make([]int, 0), food, 0}
    game.snake = append(game.snake, 0)
    return game
}


/** Moves the snake.
        @param direction - 'U' = Up, 'L' = Left, 'R' = Right, 'D' = Down 
        @return The game's score after the move. Return -1 if game over. 
        Game over when snake crosses the screen boundary or bites its body. */
func (this *SnakeGame) Move(direction string) int {
    if this.foodIndex == -1 {
        return -1
    }
    
    head := this.snake[0]
    r, c := head / this.width, head % this.width
    
    switch direction {
    case "U":
        r--
    case "L":
        c--
    case "R":
        c++
    case "D":
        r++
    }
    
    newHead := r*this.width + c
    
    tail := this.snake[len(this.snake)-1]
    this.snake = this.snake[:len(this.snake)-1]
    
    // out of boundary or eating body
    if r < 0 || r >= this.height || c < 0 || c >= this.width {
        this.foodIndex = -1
        return -1
    }
    for _, coor := range this.snake {
        if coor == newHead {
            this.foodIndex = -1
            return -1
        }
    }
    
    // append new head
    this.snake = append([]int{newHead}, this.snake...)
    
    // eating food, keeping tail
    if this.foodIndex < len(this.food) && r == this.food[this.foodIndex][0] && c == this.food[this.foodIndex][1] {
        this.snake = append(this.snake, tail)
        this.foodIndex++;
    }

    return this.foodIndex
}


/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */
