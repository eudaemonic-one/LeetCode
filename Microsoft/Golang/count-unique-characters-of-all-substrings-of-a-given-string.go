func uniqueLetterString(s string) int {
    res := 0
    index := make([][2]int, 26)
    for i := 0; i < 26; i++ {
        index[i] = [2]int{-1, -1}
    }
    n := len(s)
    mod := int(10e9 + 7)
    
    for i := 0; i < n; i++ {
        c := int(s[i]-'A')
        res = (res + (i - index[c][1]) * (index[c][1] - index[c][0]) % mod) % mod
        index[c] = [2]int{index[c][1], i}
    }
   
    for c := 0; c < 26; c++ {
        res = (res + (n - index[c][1]) * (index[c][1] - index[c][0]) % mod) % mod
    }
    
    return res % mod
}
