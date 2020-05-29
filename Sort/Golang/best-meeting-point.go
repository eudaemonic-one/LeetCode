func minTotalDistance(grid [][]int) int {
    rows := make([]int, 0)
    cols := make([]int, 0)
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                rows = append(rows, i)
                cols = append(cols, j)
            }
        }
    }
    sort.Ints(cols)
    originX, originY := rows[len(rows)/2], cols[len(cols)/2]
    return getDistance1D(rows, originX) + getDistance1D(cols, originY)
}

func getDistance1D(points []int, origin int) int {
    dist := 0
    for _, point := range points {
        dist += abs(point - origin)
    }
    return dist
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
