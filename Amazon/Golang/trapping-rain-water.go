// Approach 4: Two Pointers

func trap(height []int) int {
    res := 0
    l, r := 0, len(height)-1
    leftMax, rightMax := 0, 0
    for l < r {
        if height[l] < height[r] {
            if height[l] >= leftMax {
                leftMax = height[l]
            } else {
                res += leftMax - height[l]
            }
            l++
        } else {
            if height[r] >= rightMax {
                rightMax = height[r]
            } else {
                res += rightMax - height[r]
            }
            r--
        }
    }
    return res
}

// Approach 3: Stack

// func trap(height []int) int {
//     res := 0
//     stack := make([]int, 0)
//     for i := 0; i < len(height); i++ {
//         for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
//             top := stack[len(stack)-1]
//             stack = stack[:len(stack)-1]
//             if len(stack) == 0 {
//                 break
//             }
//             dist := i - stack[len(stack)-1] - 1
//             height := min(height[i], height[stack[len(stack)-1]]) - height[top]
//             res += dist * height
//         }
//         stack = append(stack, i)
//     }
//     return res
// }

// Approach 2: Dynamic Programming

// func trap(height []int) int {
//     if len(height) == 0 {
//         return 0
//     }
//     leftMax := make([]int, len(height))
//     leftMax[0] = height[0]
//     for l := 1; l < len(height); l++ {
//         leftMax[l] = max(leftMax[l-1], height[l])
//     } 
//     rightMax := make([]int, len(height))
//     rightMax[len(height)-1] = height[len(height)-1]
//     for r := len(height)-2; r >= 0; r-- {
//         rightMax[r] = max(rightMax[r+1], height[r])
//     }
//     res := 0
//     for i := 1; i < len(height)-1; i++ {
//         res += min(leftMax[i], rightMax[i]) - height[i]
//     }
//     return res
// }

// Approach 1: Brute Force

// func trap(height []int) int {
//     res := 0
//     for i := 0; i < len(height); i++ {
//         leftMax := 0
//         for l := i; l >= 0; l-- {
//             leftMax = max(leftMax, height[l])
//         } 
//         rightMax := 0
//         for r := i; r < len(height); r++ {
//             rightMax = max(rightMax, height[r])
//         }
//         res += min(leftMax, rightMax) - height[i]
//     }
//     return res
// }

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
