// Approach #2: Optimized Bit Manipulation with Dynamic Programming

func countBits(num int) []int {
    dp := make([]int, num+1)
    for i := 1; i <= num; i++ {
        dp[i] = dp[i&(i-1)] + 1
    }
    return dp
}

// Approach #1: Naive Bit Manipulation

// func countBits(num int) []int {
//     counts := make([]int, num+1)
//     for i := 0; i <= num; i++ {
//         ones := 0
//         x := i
//         for x > 0 {
//             x = x & (x-1)
//             ones++
//         }
//         counts[i] = ones
//     }
//     return counts
// }
