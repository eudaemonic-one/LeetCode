func nextGreaterElement(nums1 []int, nums2 []int) []int {
    table := make(map[int]int)
    stack := make([]int, 0)
    for _, num := range nums2 {
        for len(stack) > 0 && stack[len(stack)-1] < num {
            table[stack[len(stack)-1]] = num
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, num)
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
