// Approach #2: DP (Optimized)
func maximalRectangle(matrix [][]byte) int {
    m := len(matrix)
    if m == 0 {
        return 0
    }
    n := len(matrix[0])
    if n == 0 {
        return 0
    }
    
    res := 0
    dp := make([]int, n)
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                dp[j] += 1
            } else {
                dp[j] = 0
            }
        }
        if area := largestRectangleArea(dp); area > res {
            res = area
        }
    }
    
    return res
}

func largestRectangleArea(heights []int) int {
    res := 0
    stack := []int{-1}
    for i, height := range heights {
        for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= height {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if area := heights[top]*(i-stack[len(stack)-1]-1); area > res {
                res = area
            }
        }
        stack = append(stack, i)
    }
    for stack[len(stack)-1] != -1 {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if area := heights[top]*(len(heights)-stack[len(stack)-1]-1); area > res {
            res = area
        }
    }
    return res
}

// Approach #1: DP

// func maximalRectangle(matrix [][]byte) int {
//     m := len(matrix)
//     if m == 0 {
//         return 0
//     }
//     n := len(matrix[0])
//     if n == 0 {
//         return 0
//     }
//     dp := make([][]int, m)
//     for i := 0; i < m; i++ {
//         dp[i] = make([]int, n)
//     }
//     res := 0
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if matrix[i][j] == '1' {
//                 if j == 0 {
//                     dp[i][j] = 1
//                 } else {
//                     dp[i][j] = dp[i][j-1] + 1
//                 }
//                 width := dp[i][j]
//                 for k := i; k >= 0; k-- {
//                     if dp[k][j] < width {
//                         width = dp[k][j]
//                     }
//                     if area := width * (i-k+1); area > res {
//                         res = area
//                     }
//                 }
//             }
//         }
//     }
//     return res
// }
