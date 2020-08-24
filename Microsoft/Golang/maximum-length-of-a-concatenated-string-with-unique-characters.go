func maxLength(arr []string) int {
    if len(arr) == 0 {
        return 0
    }
    res := 0
    dp := []int{0}
    for _, word := range arr {
        uniqueBitmap := 0
        duplicateBitmap := 0
        for i := 0; i < len(word); i++ {
            mask := 1 << int(word[i]-'a')
            duplicateBitmap |= uniqueBitmap & mask
            if duplicateBitmap != 0 {
                break
            }
            uniqueBitmap |= mask
        }
        if duplicateBitmap != 0 {
            continue
        }
        
        for i := len(dp)-1; i >= 0; i-- {
            if dp[i] & uniqueBitmap != 0 {
                continue
            }
            newBitmap := dp[i] | uniqueBitmap
            dp = append(dp, newBitmap)
            res = max(res, bits.OnesCount(uint(newBitmap)))
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
