func findMin(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    l, r := 0, len(nums)-1
    if nums[r] >= nums[0] {
        return nums[0]
    }
    for l < r {
        m := l + (r-l)/2
        if nums[m] > nums[m+1] {
            return nums[m+1]
        }
        if nums[m] > nums[0] {
            l = m + 1
        } else {
            r = m
        }
    }
    return -1
}
