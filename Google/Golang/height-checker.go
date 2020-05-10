func heightChecker(heights []int) int {
    res := 0
    if len(heights) == 0 {
        return 0
    }
    dup := make([]int, len(heights))
    copy(dup, heights)
    sort.Ints(dup)
    for i := 0; i < len(heights); i++ {
        if heights[i] != dup[i] {
            res += 1
        }
    }
    return res
}
