func search(nums []int, target int) int {
    n := len(nums)
    if n == 0 {
        return -1
    }
    l, r := 0, n-1
    for l <= r {
        m := l + (r-l)/2
        if nums[m] == target {
            return m
        } else if nums[m] >= nums[l] {
            if nums[l] <= target && target < nums[m] {
                r = m - 1
            } else {
                l = m + 1
            }
        } else {
            if nums[m] < target && target <= nums[r] {
                l = m + 1
            } else {
                r = m - 1
            }
        }
    }
    return -1
}
