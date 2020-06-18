func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    if m == 0 {
        return false
    }
    n := len(matrix[0])
    if n == 0 {
        return false
    }
    
    l, r := 0, m*n-1
    for l <= r {
        pi := l + (r-l)/2
        val := matrix[pi/n][pi%n]
        if target == val {
            return true
        } else if val > target {
            r = pi - 1
        } else {
            l = pi + 1
        }
    }
    return false
}
