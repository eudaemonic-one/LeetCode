func lengthOfLongestSubstringKDistinct(s string, k int) int {
    res := 0
    counter := make(map[byte]int)
    l, r := 0, 0
    for r < len(s) {
        counter[s[r]]++
        r++
        
        for len(counter) > k {
            counter[s[l]]--
            if counter[s[l]] == 0 {
                delete(counter, s[l])
            }
            l++
        }
        
        if r - l > res {
            res = r - l
        }
    }
    return res
}
