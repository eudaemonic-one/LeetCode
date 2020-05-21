func maxLength(arr []string) int {
    res := 0
    dp := []int{0}
    for _, word := range arr {
        uniqueSet, duplicateSet := 0, 0
        for _, ch := range word {
            bit := 1 << int(ch - 'a')
            duplicateSet |= uniqueSet & bit
            uniqueSet |= bit
        }
        if duplicateSet > 0 {
            continue
        }
        for i := len(dp)-1; i >= 0; i-- {
            if dp[i] & uniqueSet > 0 {
                continue
            }
            newSet := dp[i] | uniqueSet
            dp = append(dp, newSet)
            res = max(res, bits.OnesCount(uint(newSet)))
        }
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
