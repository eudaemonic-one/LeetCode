func minCut(s string) int {
    n := len(s)
    cut := make([]int, n)
    dp := make([][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]bool, n)
        dp[i][i] = true
    }
    
    for i := 0; i < n; i++ {
        cut[i] = i
        for j := 0; j <= i; j++ {
            if s[i] == s[j] && (j+1 > i-1 || dp[j+1][i-1]) {
                dp[j][i] = true
                if j == 0 {
                    cut[i] = 0
                } else if cut[j-1]+1 < cut[i] {
                    cut[i] = cut[j-1]+1
                }
            }
        }
    }
    
    return cut[n-1]
}
