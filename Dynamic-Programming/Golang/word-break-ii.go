func wordBreak(s string, wordDict []string) []string {
    res := make([]string, 0)
    dict := make(map[string]bool)
    for _, word := range wordDict {
        dict[word] = true
    }
    dp := make([]bool, len(s))
    
    for j := 0; j < len(s); j++ {
        for i := 0; i <= j; i++ {
            if _, ok := dict[s[i:j+1]]; ok && (i == 0 || dp[i-1]) {
                dp[j] = true
            }
        }
    }
    
    if !dp[len(s)-1] {
        return res
    }
    
    dfs(s, dict, 0, &[]string{}, &res)
    
    return res
}

func dfs(s string, dict map[string]bool, i int, path *[]string, res *[]string) {
    if i == len(s) {
        *res = append(*res, strings.Join(*path, " "))
        return
    }
    for j := i; j < len(s); j++ {
        if _, ok := dict[s[i:j+1]]; ok {
            *path = append(*path, s[i:j+1])
            dfs(s, dict, j+1, path, res)
            *path = (*path)[:len(*path)-1]
        }
    }
}
