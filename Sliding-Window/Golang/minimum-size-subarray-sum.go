func minSubArrayLen(s int, nums []int) int {
    res := len(nums) + 1
    sum := 0
    l, r := 0, 0
    for r < len(nums) {
        sum += nums[r]
        r++
        for sum >= s {
            res = min(res, r-l)
            sum -= nums[l]
            l++
        }
    }
    
    if res == len(nums)+1 {
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
