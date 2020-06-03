func getSkyline(buildings [][]int) [][]int {
    res := make([][]int, 0)
    n := len(buildings)
    // Base case; no building
    if n == 0 {
        return res
    }
    // Base case; only one building
    if n == 1 {
        res = append(res, []int{buildings[0][0], buildings[0][2]})
        res = append(res, []int{buildings[0][1], 0})
        return res
    }
    // recursively divide the input into two subproblems
    leftSkyline := getSkyline(buildings[:n/2])
    rightSkyline := getSkyline(buildings[n/2:])
    return mergeSkylines(leftSkyline, rightSkyline)
}

func mergeSkylines(leftSkyline, rightSkyline [][]int) [][]int {
    res := make([][]int, 0)
    nL, nR := len(leftSkyline), len(rightSkyline)
    pL, pR := 0, 0
    currY, leftY, rightY, maxY := 0, 0, 0, 0
    x := 0
    for pL < nL && pR < nR {
        pointL := leftSkyline[pL]
        pointR := rightSkyline[pR]
        // pick up the smallest x
        if pointL[0] < pointR[0] {
            x = pointL[0]
            leftY = pointL[1]
            pL++
        } else {
            x = pointR[0]
            rightY = pointR[1]
            pR++
        }
        // max height (i.e. y) between both skylines
        maxY = max(leftY, rightY)
        // update output if there is a skyline change
        if currY != maxY {
            res = updateSkyline(res, x, maxY)
            currY = maxY
        }
    }
    res = appendSkyline(res, leftSkyline, pL, nL, currY)
    res = appendSkyline(res, rightSkyline, pR, nR, currY)
    return res
}

func updateSkyline(res [][]int, x, y int) [][]int {
    if len(res) == 0 || res[len(res)-1][0] != x {
        res = append(res, []int{x, y})
    } else {
        res[len(res)-1][1] = y
    }
    return res
}

func appendSkyline(res [][]int, skyline [][]int, p, n, currY int) [][]int {
    for p < n {
        point := skyline[p]
        x := point[0]
        y := point[1]
        p++
        // update skylines
        if currY != y {
            res = updateSkyline(res, x, y)
            currY = y
        }
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
