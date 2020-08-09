func nextGreaterElement(nums1 []int, nums2 []int) []int {
    table := make(map[int]int)
    stack := make([]int, 0)
    for i := len(nums2)-1; i >= 0; i-- {
        for len(stack) > 0 && nums2[i] >= stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
        }
        if len(stack) == 0 {
            table[nums2[i]] = -1
        } else {
            table[nums2[i]] = stack[len(stack)-1]
        }
        stack = append(stack, nums2[i])
    }
    res := make([]int, len(nums1))
    for i, num := range nums1 {
        if val, ok := table[num]; ok {
            res[i] = val
        } else {
            res[i] = -1
        }
    }
    return res
}
