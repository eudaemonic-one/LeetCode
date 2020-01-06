func sortColors(nums []int)  {
    i := 0
    zero, two := 0, len(nums)-1
    for i <= two {
        if nums[i] == 0 {
            nums[zero], nums[i] = nums[i], nums[zero]
            zero++
            i++
        } else if nums[i] == 2 {
            nums[two], nums[i] = nums[i], nums[two]
            two--
        } else {
            i++
        }
    }
}
