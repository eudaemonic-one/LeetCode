func maxSumDivThree(nums []int) int {
    dp := make([]int, 3)
    tmp := make([]int, 3)
    for _, v := range nums {
        copy(tmp, dp)
        for i := range dp {
            n := tmp[i] + v
            dp[n % 3] = max(dp[n % 3], n)
        }
    }
    return dp[0]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}