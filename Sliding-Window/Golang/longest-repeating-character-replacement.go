func characterReplacement(s string, k int) int {
    if len(s) == 0 {
        return 0
    }
    res := 0
    count := make([]int, 26)
    maxCount := 0
    l, r := 0, 0
    for r < len(s) {
        ch := int(s[r]-'A')
        count[ch]++
        maxCount = max(maxCount, count[ch])
        r++
        
        for r - l - maxCount > k {
            count[int(s[l]-'A')]--
            l++
        }
        
        res = max(res, r-l)
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
