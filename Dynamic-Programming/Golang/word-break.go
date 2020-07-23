func wordBreak(s string, wordDict []string) bool {
    if len(s) == 0 || len(wordDict) == 0 {
        return false
    }
    
    wordSet := make(map[string]bool)
    for _, word := range wordDict {
        wordSet[word] = true
    }
    dp := make([]bool, len(s))
    
    for j := 0; j < len(s); j++ {
        for i := 0; i <= j; i++ {
            if _, ok := wordSet[s[i:j+1]]; (i == 0 || dp[i-1]) && ok {
                dp[j] = true
            }
        }
    }
    
    return dp[len(s)-1]
}
