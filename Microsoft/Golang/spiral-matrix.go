func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }
    
    res := make([]int, 0)
    left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1
    
    for left <= right && up <= down {
        for i := left; i <= right; i++ {
            res = append(res, matrix[up][i])
        }
        for j := up+1; j <= down; j++ {
            res = append(res, matrix[j][right])
        }
        if left < right && up < down {
            for i := right-1; i > left; i-- {
                res = append(res, matrix[down][i])
            }
            for j := down; j > up; j-- {
                res = append(res, matrix[j][left])
            }
        }
        up++
        down--
        right--
        left++
    }
    
    return res
}
