func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
    res := 0
    finished := false
    queue := make([]int, 0)
    for i := 0; i < len(initialBoxes); i++ {
        queue = append(queue, initialBoxes[i])
    }
    for len(queue) > 0 {
        l := len(queue)
        finished = true
        for j := 0; j < l; j++ {
            box := queue[0]
            queue = queue[1:]
            if status[box] == 1 { // if we can open the box
                finished = false
                res += candies[box]
                for _, key := range keys[box] {
                    status[key] = 1
                }
                for _, p := range containedBoxes[box] {
                    queue = append(queue, p)
                }
            } else {
                queue = append(queue, box) // try this box in next cycle
            }
        }
        if finished {
            return res
        }
    }
    return res
}
