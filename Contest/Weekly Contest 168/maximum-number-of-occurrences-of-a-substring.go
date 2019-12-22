func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
    if len(s) < minSize {
        return 0
    }
    res := 0
    uniqueCharMap := make(map[byte]int)
    substrsMap := make(map[string]int)
    for l, r := 0, 0; r < len(s); r++ {
        uniqueCharMap[s[r]]++
        if r-l+1 > minSize {
            uniqueCharMap[s[l]]--
            if uniqueCharMap[s[l]] == 0 {
                delete(uniqueCharMap, s[l])
            }
            l++
        }
        if minSize <= r-l+1 && r-l+1 <= maxSize && len(uniqueCharMap) <= maxLetters {
            substrsMap[s[l:r+1]]++
        }
    }
    for _, v := range substrsMap {
        if v > res {
            res = v
        }
    }
    return res
}
