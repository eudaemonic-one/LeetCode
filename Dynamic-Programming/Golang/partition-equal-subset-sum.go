// Approach #2: DP with tricks

func canPartition(nums []int) bool {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if sum & 1 == 1 {
        return false
    }
    
    dp := make([]bool, sum/2+1)
    dp[0] = true
    
    for i := 1; i < len(nums)+1; i++ {
        for j := sum/2; j > 0; j-- {
            if j >= nums[i-1] {
                dp[j] = dp[j] || dp[j-nums[i-1]]
            }
        }
    }
    
    return dp[sum/2]
}

// Approach #1: DP (0-1 Knapsack)

// func canPartition(nums []int) bool {
//     sum := 0
//     for _, num := range nums {
//         sum += num
//     }
//     if sum & 1 == 1 {
//         return false
//     }
    
//     dp := make([][]bool, len(nums)+1)
//     for i := 0; i < len(nums)+1; i++ {
//         dp[i] = make([]bool, sum/2+1)
//         dp[i][0] = true
//     }
    
//     for i := 1; i < len(nums)+1; i++ {
//         for j := 1; j < sum/2+1; j++ {
//             if j < nums[i-1] {
//                 dp[i][j] = dp[i-1][j]
//             } else {
//                 dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
//             }
//         }
//     }
    
//     return dp[len(nums)][sum/2]
// }
