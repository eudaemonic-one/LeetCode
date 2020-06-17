func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    if m == 0 {
        return false
    }
    n := len(matrix[0])
    if n == 0 {
        return false
    }
    lo, pi, hi := 0, 0, 0
    for i := 0; i < min(m, n); i++ {
        lo, hi = i, n-1
        for lo <= hi {
            pi = lo + (hi-lo)/2
            if matrix[i][pi] == target {
                return true
            } else if matrix[i][pi] < target {
                lo = pi + 1
            } else {
                hi = pi - 1
            }
        }
        lo, hi = i, m-1
        for lo <= hi {
            pi = lo + (hi-lo)/2
            if matrix[pi][i] == target {
                return true
            } else if matrix[pi][i] < target {
                lo = pi + 1
            } else {
                hi = pi - 1
            }
        }
    }
    return false
}
                       
func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
