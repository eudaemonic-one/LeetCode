func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    return max(robBetween(nums, 0, len(nums)-1), robBetween(nums, 1, len(nums)))
}

func robBetween(nums []int, start, end int) int {
    a, b, c := 0, 0, 0
    for i := end-1; i >= start; i-- {
        a = max(b, nums[i]+c)
        c = b
        b = a
    }
    return a
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
