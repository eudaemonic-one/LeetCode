func rotate(matrix [][]int)  {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return
    }
    m, n := len(matrix), len(matrix[0])
    // flip the image by the diagonal from top left to bottom right
    for i := 0; i < m; i++ {
        for j := i+1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    // flip the image horizontally
    for i := 0; i < m; i++ {
        for j := 0; j < n/2; j++ {
            matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
        }
    }
}
