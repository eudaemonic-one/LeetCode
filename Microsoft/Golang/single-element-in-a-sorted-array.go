func singleNonDuplicate(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    l, r := 0, len(nums)-1
    for l < r {
        m := l + (r-l)/2
        halvesAreEven := (r - m) % 2 == 0
        if nums[m] == nums[m+1] {
            if halvesAreEven {
                l = m + 2
            } else {
                r = m - 1
            }
        } else if nums[m-1] == nums[m] {
            if halvesAreEven {
                r = m - 2
            } else {
                l = m + 1
            }
        } else {
            return nums[m]
        }
    }
    return nums[l]
}
