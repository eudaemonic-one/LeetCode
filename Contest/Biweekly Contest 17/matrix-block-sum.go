func matrixBlockSum(mat [][]int, K int) [][]int {
    if len(mat) == 0 || len(mat[0]) == 0 {
        return mat
    }
    m, n := len(mat), len(mat[0])
    answer := make([][]int, m)
    for i := 0; i < m; i++ {
        answer[i] = make([]int, n)
        for j := 0; j < n; j++ {
            sum := 0
            R, C := i - K, j - K
            if R < 0 {
                R = 0
            }
            if C < 0 {
                C = 0
            }
            for r := R; r <= i+K && r < m; r++ {
                for c := C; c <= j+K && c < n; c++ {
                    sum += mat[r][c]
                }
            }
            answer[i][j] = sum
        }
    }
    return answer
}