func prisonAfterNDays(cells []int, N int) []int {
    if len(cells) == 0 {
        return cells
    }
    
    isFastForward := false
    visited := make(map[int]int)
    
    for N > 0 {
        if !isFastForward {
            stateBitmap := cellsToBitmap(cells)
            if lastTime, ok := visited[stateBitmap]; ok {
                N %= lastTime - N
                isFastForward = true
            } else {
                visited[stateBitmap] = N
            }
        }
        
        if N > 0 {
            N -= 1
            cells = nextState(cells)
        }
    }
    
    return cells
}

func cellsToBitmap(cells []int) int {
    bitmap := 0
    for _, cell := range cells {
        bitmap <<= 1
        bitmap = bitmap | cell
    }
    return bitmap
}

func nextState(state []int) []int {
    newState := make([]int, len(state))
    newState[0] = 0
    for i := 1; i < len(state)-1; i++ {
        prev, next := state[i-1], state[i+1]
        if prev == next {
            newState[i] = 1
        } else {
            newState[i] = 0
        }
    }
    newState[len(newState)-1] = 0
    return newState
}
