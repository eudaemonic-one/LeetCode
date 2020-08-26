func longestDupSubstring(S string) string {
    res := ""
    l, r := 1, len(S)-1
    for l <= r {
        m := l + (r-l)/2
        isValid, subStr := search(S, m)
        if isValid {
            res = subStr
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return res
}

func search(str string, length int) (bool, string) {
    if len(str) < length {
        return false, ""
    }
    seen := make(map[string]bool)
    for i := 0; i < len(str)-length+1; i++ {
        subStr := str[i:i+length]
        if _, ok := seen[subStr]; ok {
            return true, subStr
        }
        seen[subStr] = true
    }
    return false, ""
}
