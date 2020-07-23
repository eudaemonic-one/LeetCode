func longestOnes(A []int, K int) int {
    res := 0
    zeros, ones := 0, 0
    l, r := 0, 0
    for r < len(A) {
        if A[r] == 0 {
            zeros++
        } else {
            ones++
        }
        r++
        
        for zeros > K {
            if A[l] == 0 {
                zeros--
            } else {
                ones--
            }
            l++
        }
        
        if zeros+ones > res {
            res = zeros + ones
        }
    }
    return res
}
