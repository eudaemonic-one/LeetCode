// Approach 3: Search Space Reduction

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    r, c := len(matrix)-1, 0
    for r >= 0 && c < len(matrix[0]) {
        if matrix[r][c] > target {
            r--
        } else if matrix[r][c] < target {
            c++
        } else {
            return true
        }
    }
    return false
}

// Approach 2: Divide and Conquer

// func searchMatrix(matrix [][]int, target int) bool {
//     if len(matrix) == 0 || len(matrix[0]) == 0 {
//         return false
//     }
//     return searchHelper(matrix, target, 0, 0, len(matrix[0])-1, len(matrix)-1)
// }

// func searchHelper(matrix [][]int, target, left, up, right, down int) bool {
//     if left > right || up > down || target < matrix[up][left] || target > matrix[down][right] {
//         return false
//     }
//     mid := left + (right-left)/2
//     row := up
//     for row <= down && matrix[row][mid] <= target {
//         if matrix[row][mid] == target {
//             return true
//         }
//         row++
//     }
//     return searchHelper(matrix, target, left, row, mid-1, down) || searchHelper(matrix, target, mid+1, up, right, row-1)
// }

// Approach 1: Binary Search

// func searchMatrix(matrix [][]int, target int) bool {
//     m := len(matrix)
//     if m == 0 {
//         return false
//     }
//     n := len(matrix[0])
//     if n == 0 {
//         return false
//     }
//     lo, pi, hi := 0, 0, 0
//     for i := 0; i < min(m, n); i++ {
//         lo, hi = i, n-1
//         for lo <= hi {
//             pi = lo + (hi-lo)/2
//             if matrix[i][pi] == target {
//                 return true
//             } else if matrix[i][pi] < target {
//                 lo = pi + 1
//             } else {
//                 hi = pi - 1
//             }
//         }
//         lo, hi = i, m-1
//         for lo <= hi {
//             pi = lo + (hi-lo)/2
//             if matrix[pi][i] == target {
//                 return true
//             } else if matrix[pi][i] < target {
//                 lo = pi + 1
//             } else {
//                 hi = pi - 1
//             }
//         }
//     }
//     return false
// }
                       
// func min(x, y int) int {
//     if x < y {
//         return x
//     }
//     return y
// }
