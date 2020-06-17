func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    l, r := matrix[0][0], matrix[n-1][n-1]
    for l < r {
        m := l + (r-l)/2
        if cnt := countSmaller(matrix, n, m); cnt < k {
            l = m + 1
        } else {
            r = m
        }
    }
    return l
}

func countSmaller(matrix [][]int, n, target int) int {
    count := 0
    j := n - 1
    for i := 0; i < n; i++ {
        for j >= 0 && matrix[i][j] > target {
            j--
        }
        if j < 0 {
            break
        }
        count += j + 1
    }
    return count
}
