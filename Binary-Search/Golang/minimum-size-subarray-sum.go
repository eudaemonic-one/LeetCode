func minSubArrayLen(s int, nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    res := math.MaxInt64
    sums := make([]int, n+1)
    for i := 1; i <= n; i++ {
        sums[i] = sums[i-1] + nums[i-1]
    }
    for i := 1; i <= n; i++ {
        target := s + sums[i-1]
        l, r := i, len(sums)
        for l < r {
            m := l + (r-l)/2
            if sums[m] < target {
                l = m + 1
            } else {
                r = m
            }
        }
        if l < len(sums) {
            res = min(res, l-i+1)
        }
    }
    if res == math.MaxInt64 {
        return 0
    }
    return res
}

func min (x, y int) int {
    if x < y {
        return x
    }
    return y
}
