func lengthOfLongestSubstring(s string) int {
    l, r, res := 0, 0, 0
    if len(s) == 0 {
        return res
    }
    dict := make(map[byte]int)
    for r < len(s) && l <= r {
        ch := s[r]
        if _, ok := dict[ch]; !ok {
            dict[ch] = 1
        } else {
            dict[ch] += 1
        }
        for dict[ch] > 1 {
            dict[s[l]] -= 1
            l++
        }
        r++
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
