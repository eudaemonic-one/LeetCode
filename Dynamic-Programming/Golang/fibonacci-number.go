func fib(N int) int {
    if N <= 1 {
        return N
    }
    if N == 2 {
        return 1
    }
    a, b, c := 1, 1, 0
    for i := 3; i <= N; i++ {
        c = a + b
        b = a
        a = c
    }
    return c
}

// Approach #4: DP with Tricks

// func fib(N int) int {
//     if N <= 1 {
//         return N
//     }
//     if N == 2 {
//         return 1
//     }
//     a, b, c := 1, 1, 0
//     for i := 3; i <= N; i++ {
//         c = a + b
//         b = a
//         a = c
//     }
//     return c
// }

// Approach #3: Bottom-up DP (Iterative)

// func fib(N int) int {
//     if N <= 1 {
//         return N
//     }
//     dp := make([]int, N+1)
//     dp[1] = 1
//     for i := 2; i <= N; i++ {
//         dp[i] = dp[i-1] + dp[i-2]
//     }
//     return dp[N]
// }

// Approach #2: Recursion with Memo

// var memo = map[int]int{0:0, 1:1}

// func fib(N int) int {
//     if N <= 1 {
//         return N
//     }
//     return helper(N)
// }

// func helper(N int) int {
//     if val, ok := memo[N]; ok {
//         return val
//     }
//     memo[N] = helper(N-1) + helper(N-2)
//     return memo[N]
// }

// Approach #1: Recursion

// func fib(N int) int {
//     if N <= 1 {
//         return N
//     }
//     return fib(N-1) + fib(N-2)
// }

