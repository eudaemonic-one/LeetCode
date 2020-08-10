// Approach 3: Expand Around Center

func longestPalindrome(s string) string {
    if len(s) <= 1 {
        return s
    }
    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        len1 := expandAroundCenter(s, i, i)
        len2 := expandAroundCenter(s, i, i+1)
        length := max(len1, len2)
        if length > end-start {
            start = i - (length-1)/2
            end = i + length/2
        }
    }
    return s[start:end+1]
}

func expandAroundCenter(s string, l int, r int) int {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }
    return r - l - 1
}

func max(x int, y int) int {
    if x > y {
        return x
    }
    return y
}

// Approach 2: Dynamic Programming

// func longestPalindrome(s string) string {
//     if len(s) <= 1 {
//         return s
//     }

//     res := string(s[0])
//     n := len(s)
//     // dp[i][j]: whether the substring from i to j is a palindrome substring
//     dp := make([][]bool, n)
//     for i := 0; i < n; i++ {
//         dp[i] = make([]bool, n)
//         dp[i][i] = true
//     }
        
//     for i := n-2; i >= 0; i-- {
//         for j := i+1; j < n; j++ {
//             if s[i] == s[j] {
//                 dp[i][j] = j-i <= 1 || dp[i+1][j-1]
//                 if dp[i][j] && j-i+1 > len(res) {
//                     res = s[i:j+1]
//                 }
//             }
//         }
//     }
    
//     return res
// }

// Approach 1: Brute Force
