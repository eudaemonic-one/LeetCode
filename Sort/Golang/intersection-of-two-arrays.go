func intersection(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    res := make([]int, 0)
    l, r := 0, 0
    for l < len(nums1) && r < len(nums2) {
        left, right := nums1[l], nums2[r]
        if left == right {
            res = append(res, left)
            for l < len(nums1) && left == nums1[l] {
                l++
            }
            for r < len(nums2) && right == nums2[r] {
                r++
            }
        } else if left < right {
            for l < len(nums1) && left == nums1[l] {
                l++
            }
        } else {
            for r < len(nums2) && right == nums2[r] {
                r++
            }
        }
    }
    return res
}
