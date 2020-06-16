func searchRange(nums []int, target int) []int {
    l, r := 0, len(nums)
    for l < r {
        m := l + (r-l)/2
        if nums[m] >= target {
            r = m
        } else {
            l = m + 1
        }
    }
    if l >= len(nums) || nums[l] != target {
        return []int{-1,-1}
    }
    left := l
    r = len(nums)
    for l < r {
        m := l + (r-l)/2
        if nums[m] > target {
            r = m
        } else {
            l = m + 1
        }
    }
    right := l - 1
    return []int{left,right}
}
