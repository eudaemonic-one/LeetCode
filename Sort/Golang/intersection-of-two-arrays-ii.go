func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    res := make([]int, 0)
    l, r := 0, 0
    for l < len(nums1) && r < len(nums2) {
        left, right := nums1[l], nums2[r]
        if left == right {
            res = append(res, left)
            l++
            r++
        } else if left < right {
            for l < len(nums1) && nums1[l] < right {
                l++
            }
        } else {
            for r < len(nums2) && nums2[r] < left {
                r++
            }
        }
    }
    return res
}
