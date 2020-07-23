func longestPalindrome(s string) string {
    if len(s) <= 1 {
        return s
    }
    res := string(s[0])
    n := len(s)
    dp := make([][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]bool, n)
        dp[i][i] = true
    }
    
    for i := n-1; i >= 0; i-- {
        for j := i+1; j < n; j++ {
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1] || j-i < 3
                if dp[i][j] && j-i+1 > len(res) {
                    res = s[i:j+1]
                }
            }
        }
    }
    
    return res
}
