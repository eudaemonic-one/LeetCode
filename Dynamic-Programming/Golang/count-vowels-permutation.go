// Approach #3: Bottom-up DP

var graph = [][]int{[]int{1,2,4}, []int{0,2}, []int{1,3}, []int{2}, []int{2,3}}

func countVowelPermutation(n int) int {
    dp := make([][]int, 2)
    for i := 0; i < 2; i++ {
        dp[i] = make([]int, 5)
    }
    for j := 0; j < 5; j++ {
        dp[0][j] = 1
    }
    
    for i := 1; i < n; i++ {
        for j := 0; j < 5; j++ {
            dp[i%2][j] = 0
            for _, v := range graph[j] {
                dp[i%2][j] = (dp[i%2][j] + dp[(i-1)%2][v]) % (1e09+7)
            }
        }
    }
    
    res := 0
    for i := 0; i < 5; i++ {
        res = (res + dp[(n-1)%2][i]) % (1e09+7)
    }
    return res
}

// Approach #2: Recursion with Memo

// var graph = [][]int{[]int{1}, []int{0,2}, []int{0,1,3,4}, []int{2,4}, []int{0}}

// func countVowelPermutation(n int) int {
//     res := 0
//     memo := make([][]int, n+1)
//     for i := 0; i < n+1; i++ {
//         memo[i] = make([]int, 5)
//         for j := 0; j < 5; j++ {
//             memo[i][j] = -1
//         }
//     }
//     for i := 0; i < 5; i++ {
//         res = (res + helper(i, n, memo)) % (1e09+7)
//     }
//     return res % (1e09+7)
// }

// func helper(ch, n int, memo [][]int) int {
//     if n == 1 {
//         return 1
//     }
//     if memo[n][ch] != -1 {
//         return memo[n][ch]
//     }
//     res := 0
//     for _, v := range graph[ch] {
//         res = (res + helper(v, n-1, memo)) % (1e09+7)
//     }
//     memo[n][ch] = res
//     return res
// }

// Approach #1: Recursion (with Backtracking)

// func countVowelPermutation(n int) int {
//     res := 0
//     helper(0, n, '_', &res)
//     return res % (1e09+7)
// }

// func helper(i, n int, lastCh byte, res *int) {
//     if i == n {
//         *res += 1
//         return
//     }
//     for _, ch := range []byte("aeiou") {
//         if lastCh != '_' {
//             if (lastCh == 'a' && ch != 'e') ||
//                     (lastCh == 'e' && (ch != 'a' && ch != 'i'))  ||
//                     (lastCh == 'i' && ch == 'i') ||
//                     (lastCh == 'o' && (ch != 'i' && ch != 'u')) ||
//                     (lastCh == 'u' && ch != 'a') {
//                 continue
//             }
//         }
//         helper(i+1, n, ch, res)
//     }
// }
