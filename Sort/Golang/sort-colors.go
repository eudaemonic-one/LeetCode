func sortColors(nums []int)  {
    i, l, r := 0, 0, len(nums)-1
    for i <= r {
        n := nums[i]
        if n == 0 {
            nums[l], nums[i] = nums[i], nums[l]
            l++
            i++
        } else if n == 2 {
            nums[i], nums[r] = nums[r], nums[i]
            r--
        } else {
            i++
        }
    }
}
