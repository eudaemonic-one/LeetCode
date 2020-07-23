func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    res := nums[0]
    sum := res
    for i := 1; i < len(nums); i++ {
        sum = max(nums[i], sum+nums[i])
        res = max(res, sum)
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
