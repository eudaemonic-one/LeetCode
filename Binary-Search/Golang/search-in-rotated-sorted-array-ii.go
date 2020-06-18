func search(nums []int, target int) bool {
    if len(nums) == 0 {
        return false
    }
    l, r := 0, len(nums)-1
    for l <= r {
        m := l + (r-l)/2
        if nums[m] == target {
            return true
        } else if nums[m] > nums[l] {
            if nums[l] <= target && target < nums[m] {
                r = m - 1
            } else {
                l = m + 1
            }
        } else if nums[m] < nums[l] {
            if nums[m] < target && target <= nums[r] {
                l = m + 1
            } else {
                r = m - 1
            }
        } else {
            l++
        }
    }
    return false
}
