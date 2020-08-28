func isMatch(s string, p string) bool {
    if s == p || p == "*" {
        return true
    }
    sLen, pLen := len(s), len(p)
    if sLen == 0 || pLen == 0 {
        return false
    }
    
    dp := make([][]bool, pLen+1)
    for i := 0; i < pLen+1; i++ {
        dp[i] = make([]bool, sLen+1)
    }
    dp[0][0] = true
    
    for i := 0; i < pLen; i++ {
        if p[i] == '*' {
            j := 0
            for !dp[i][j] && j < sLen {
                j++
            }
            dp[i+1][j] = dp[i][j]
            for j < sLen {
                dp[i+1][j+1] = true
                j++
            }
        } else if p[i] == '?' {
            for j := 0; j < sLen; j++ {
                dp[i+1][j+1] = dp[i][j]
            }
        } else {
            for j := 0; j < sLen; j++ {
                dp[i+1][j+1] = dp[i][j] && p[i] == s[j]
            }
        }
    }
    
    return dp[pLen][sLen]
}
