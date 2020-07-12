// Approach #4: Greedy

func canJump(nums []int) bool {
    lastPos := len(nums)-1
    for i := len(nums)-2; i >= 0; i-- {
        if i+nums[i] >= lastPos {
            lastPos = i
        }
    }
    return lastPos == 0
}

// Approach #3: Bottom-up DP

// func canJump(nums []int) bool {
//     dp := make([]bool, len(nums))
//     dp[len(nums)-1] = true
//     for i := len(nums)-2; i >= 0; i-- {
//         maxPos := min(i+nums[i], len(nums)-1)
//         for j := i+1; j <= maxPos; j++ {
//             if dp[j] == true {
//                 dp[i] = true
//                 break
//             }
//         }
//     }
//     return dp[0]
// }

// func min(x, y int) int {
//     if x < y {
//         return x
//     }
//     return y
// }

// Approach #2: Recursion with Memo

// func canJump(nums []int) bool {
//     memo := make(map[int]bool)
//     return helper(nums, memo, 0)
// }

// func helper(nums []int, memo map[int]bool, pos int) bool {
//     if pos >= len(nums)-1 {
//         return true
//     }
//     if isValid, ok := memo[pos]; ok {
//         if isValid {
//             return true
//         }
//         return false
//     }
//     for nextPos := pos+1; nextPos <= pos+nums[pos]; nextPos++ {
//         if helper(nums, memo, nextPos) {
//             memo[pos] = true
//             return true
//         }
//     }
//     memo[pos] = false
//     return false
// }

// Approach #1: Recursion (with Backtracking)

// func canJump(nums []int) bool {
//     return helper(nums, 0)
// }

// func helper(nums []int, pos int) bool {
//     if pos >= len(nums)-1 {
//         return true
//     }
//     for nextPos := pos+1; nextPos <= pos+nums[pos]; nextPos++ {
//         if helper(nums, nextPos) {
//             return true
//         }
//     }
//     return false
// }
