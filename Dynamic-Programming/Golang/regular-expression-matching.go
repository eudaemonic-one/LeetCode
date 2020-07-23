// Approach #3: Bottom-up DP

func isMatch(s string, p string) bool {
    dp := make([][]bool, len(s)+1)
    for i := 0; i < len(s)+1; i++ {
        dp[i] = make([]bool, len(p)+1)
    }
    dp[len(s)][len(p)] = true
    
    for i := len(s); i >= 0; i-- {
        for j := len(p)-1; j >= 0; j-- {
            firstMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')
            if j+1 < len(p) && p[j+1] == '*'{
                dp[i][j] = dp[i][j+2] || (firstMatch && dp[i+1][j])
            } else {
                dp[i][j] = firstMatch && dp[i+1][j+1]
            }
        }
    }
    
    return dp[0][0]
}

// Approach #2: Recursion with Memo

// const (
//     empty int = iota
//     invalid
//     valid
// )

// func isMatch(s string, p string) bool {
//     memo := make([][]int, len(s)+1)
//     for i := 0; i < len(s)+1; i++ {
//         memo[i] = make([]int, len(p)+1)
//     }
//     return helper(memo, 0, 0, s, p)
// }

// func helper(memo [][]int, i int, j int, s string, p string) bool {
//     if memo[i][j] != empty {
//         return memo[i][j] == valid
//     }
//     res := false
//     if j == len(p) {
//         if i == len(s) {
//             res = true
//         }
//     } else {
//         firstMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')
//         if j+1 < len(p) && p[j+1] == '*' {
//             res = helper(memo, i, j+2, s, p) || (firstMatch && helper(memo, i+1, j, s, p))
//         } else {
//             res = firstMatch && helper(memo, i+1, j+1, s, p)
//         }
//     }
//     if res {
//         memo[i][j] = valid
//     } else {
//         memo[i][j] = invalid
//     }
//     return res
// }

// Approach #1: Recursion

// func isMatch(s string, p string) bool {
//     if len(p) == 0 {
//         return len(s) == 0
//     }
//     firstMatch := len(s) > 0 && (p[0] == s[0] || p[0] == '.')
    
//     if len(p) >= 2 && p[1] == '*' {
//         return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
//     } else {
//         return firstMatch && isMatch(s[1:], p[1:])
//     }
// }
