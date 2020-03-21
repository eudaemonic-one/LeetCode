func lengthOfLongestSubstring(s string) int {
    l, res := 0, 0
    if len(s) == 0 {
        return res
    }
    dict := make(map[rune]int)
    for i, c := range s {
        if _, ok := dict[c]; ok && l <= dict[c] {
            l = dict[c] + 1
        } else {
            res = max(res, i-l+1)
        }
        dict[c] = i
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
