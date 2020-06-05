func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    return searchHelper(matrix, target, 0, 0, len(matrix[0])-1, len(matrix)-1)
}

func searchHelper(matrix [][]int, target, left, up, right, down int) bool {
    if left > right || up > down {
        return false
    } else if target < matrix[up][left] || target > matrix[down][right] {
        return false
    }
    mid := left + (right-left)/2
    row := up
    for row <= down && matrix[row][mid] <= target {
        if matrix[row][mid] == target {
            return true
        }
        row++
    }
    return searchHelper(matrix, target, left, row, mid-1, down) || searchHelper(matrix, target, mid+1, up, right, row-1)
}
