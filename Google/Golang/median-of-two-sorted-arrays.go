func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums := make([]int, 0)
    nums = append(nums, nums1...)
    nums = append(nums, nums2...)
    sort.Ints(nums)
    if len(nums) & 1 == 1 {
        return float64(nums[len(nums)/2])
    }
    return float64(nums[len(nums)/2-1] + nums[len(nums)/2]) / 2
}
