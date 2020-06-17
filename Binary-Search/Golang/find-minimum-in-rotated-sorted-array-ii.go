func findMin(nums []int) int {
    if len(nums) == 0 {
        return -1
    } else if len(nums) == 1 {
        return nums[0]
    }
    l, r := 0, len(nums)-1
    if nums[r] > nums[0] {
        return nums[0]
    }
    for l < r {
        m := l + (r-l)/2
        if nums[m] > nums[r] {
            l = m + 1
        } else if nums[m] < nums[r] {
            r = m
        } else {
            r--
        }
    }
    return nums[l]
}
