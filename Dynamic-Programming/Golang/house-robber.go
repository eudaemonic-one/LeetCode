// Approach #4: DP with final tricks

func rob(nums []int) int {
    a, b, c := 0, 0, 0
    for i := len(nums)-1; i >= 0; i-- {
        a = max(b, nums[i]+c)
        c = b
        b = a
    }
    return a
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

// Approach #3: Bottom-up DP

// func rob(nums []int) int {
//     dp := make([]int, len(nums)+2)
//     for i := len(nums)-1; i >= 0; i-- {
//         dp[i] = max(dp[i+1], nums[i]+dp[i+2])
//     }
//     return dp[0]
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }

// Approach #2: Recursion with Memo

// func rob(nums []int) int {
//     memo := make(map[int]int)
//     return helper(memo, nums, 0)
// }

// func helper(memo map[int]int, nums []int, i int) int {
//     if i >= len(nums) {
//         return 0
//     }
//     if v, ok := memo[i]; ok {
//         return v
//     }
//     res := max(helper(memo, nums, i+1), nums[i]+helper(memo, nums, i+2))
//     memo[i] = res
//     return res
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }

// Approach #1: Recursion (with Backtracking)

// func rob(nums []int) int {
//     return helper(nums, 0)
// }

// func helper(nums []int, i int) int {
//     if i >= len(nums) {
//         return 0
//     }
//     res := max(helper(nums, i+1), nums[i]+helper(nums, i+2))
//     return res
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }
