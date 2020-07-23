// Approach #3: Bottom-up DP

func maxA(N int) int {
    if N <= 3 {
        return N
    } 
    dp := make([]int, N+1)
    
    for i := 1; i <= N; i++ {
        dp[i] = dp[i-1] + 1
        for j := 2; j < i; j++ {
            if a := dp[j-2]*(i-j+1); a > dp[i] {
                dp[i] = a
            }
        }
    }
    
    return dp[N]
}

// Approach #2: Recursion with Memo

// func maxA(N int) int {
//     if N <= 3 {
//         return N
//     }
//     memo := make(map[string]int)
//     return helper(memo, N, 0, 0)
// }

// func helper(memo map[string]int, n, screen, buffer int) int {
//     if n <= 0 {
//         return screen
//     }
//     key := fmt.Sprintf("%d,%d,%d", n, screen, buffer)
//     if val, ok := memo[key]; ok {
//         return val
//     }
//     res := 0
//     res = max(res, helper(memo, n-1, screen+1, buffer))         // Key 1: (A)
//     res = max(res, helper(memo, n-2, screen, screen))           // Key 2: (Ctrl-A) + Key 3: (Ctrl-C)
//     res = max(res, helper(memo, n-1, screen+buffer, buffer))    // Key 4: (Ctrl-V)
//     memo[key] = res
//     return res
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }

// Approach #1: Recursion

// func maxA(N int) int {
//     if N <= 3 {
//         return N
//     }
//     return helper(N, 0, 0)
// }

// func helper(n, screen, buffer int) int {
//     if n <= 0 {
//         return screen
//     }
//     res := 0
//     res = max(res, helper(n-1, screen+1, buffer))         // Key 1: (A)
//     res = max(res, helper(n-2, screen, screen))           // Key 2: (Ctrl-A) + Key 3: (Ctrl-C)
//     res = max(res, helper(n-1, screen+buffer, buffer))    // Key 4: (Ctrl-V)
//     return res
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }
