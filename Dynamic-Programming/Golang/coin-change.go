// Approach #3: Bottom-up DP

func coinChange(coins []int, amount int) int {
    // dp[i]: the number of coins upto the total value i
    dp := make([]int, amount+1)
    for i := 1; i < amount+1; i++ {
        dp[i] = amount+1
    }
    for i := 0; i < amount+1; i++ {
        for _, coin := range coins {
            if i-coin < 0 {
                continue
            }
            dp[i] = min(dp[i], 1+dp[i-coin])
        }
    }
    // edge case check
    if dp[amount] == amount+1 {
        return -1
    }
    return dp[amount]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

// Approach #2: Recursion with Memo

// var memo = map[int]int{}

// func coinChange(coins []int, amount int) int {
//     memo := make(map[int]int)
//     return helper(coins, memo, amount)
// }

// func helper(coins []int, memo map[int]int, target int) int {
//     if num, ok := memo[target]; ok {
//         return num
//     }
//     if target == 0 {
//         return 0
//     }
//     if target < 0 {
//         return -1
//     }
//     res := math.MaxInt64
//     for _, coin := range coins {
//         if n := helper(coins, memo, target-coin); n != -1 {
//             if 1+n < res {
//                 res = 1+n
//             }
//         }
//     }
//     if res == math.MaxInt64 {
//         memo[target] = -1
//         return -1
//     }
//     memo[target] = res
//     return res
// }

// Approach #1: Recursion (TLE)

// func coinChange(coins []int, amount int) int {
//     return helper(coins, amount)
// }

// func helper(coins []int, target int) int {
//     if target == 0 {
//         return 0
//     }
//     if target < 0 {
//         return -1
//     }
//     res := math.MaxInt64
//     for _, coin := range coins {
//         if num := helper(coins, target-coin); num != -1 {
//             if 1+num < res {
//                 res = 1+num
//             }
//         }
//     }
//     if res == math.MaxInt64 {
//         return -1
//     }
//     return res
// }
