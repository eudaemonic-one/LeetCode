func getSkyline(buildings [][]int) [][]int {
    res := make([][]int, 0)
    n := len(buildings)
    if n == 0 {
        return res
    }
    if n == 1 {
        res = append(res, []int{buildings[0][0], buildings[0][2]})
        res = append(res, []int{buildings[0][1], 0})
        return res
    }
    leftSkyline := getSkyline(buildings[:n/2])
    rightSkyline := getSkyline(buildings[n/2:])
    return mergeSkyline(leftSkyline, rightSkyline)
}

func mergeSkyline(leftSkyline, rightSkyline [][]int) [][]int {
    res := make([][]int, 0)
    pL, pR, nL, nR := 0, 0, len(leftSkyline), len(rightSkyline)
    x, currY, leftY, rightY := 0, 0, 0, 0
    for pL < nL && pR < nR {
        pointL := leftSkyline[pL]
        pointR := rightSkyline[pR]
        if pointL[0] < pointR[0] {
            x = pointL[0]
            leftY = pointL[1]
            pL++
        } else {
            x = pointR[0]
            rightY = pointR[1]
            pR++
        }
        maxY := max(leftY, rightY)
        if maxY != currY {
            res = updateSkyline(res, x, maxY)
            currY = maxY
        }
    }
    res = appendSkyline(res, leftSkyline, pL, nL, currY)
    res = appendSkyline(res, rightSkyline, pR, nR, currY)
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func updateSkyline(res [][]int, x int, y int) [][]int {
    if len(res) == 0 || res[len(res)-1][0] != x {
        res = append(res, []int{x, y})
    } else {
        res[len(res)-1][1] = y
    }
    return res
}

func appendSkyline(res, skyline [][]int, p, n, currY int) [][]int {
    for p < n {
        point := skyline[p]
        x, y := point[0], point[1]
        p++
        if y != currY {
            res = updateSkyline(res, x, y)
            currY = y
        }
    }
    return res
}
