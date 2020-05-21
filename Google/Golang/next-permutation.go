func nextPermutation(nums []int)  {
    i := len(nums) - 1
    for i > 0 && nums[i] <= nums[i-1] {
        i--
    }
    
    if i - 1 >= 0 {
        j := len(nums) - 1
        for j > 0 && nums[j] <= nums[i-1] {
            j--
        }
        nums[i-1], nums[j] = nums[j], nums[i-1]
    }
    
    l, r := i, len(nums) - 1
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}
