func twoSum(nums []int, target int) []int {
    set := make(map[int]int)
    for i, n := range nums {
        if idx, ok := set[n]; ok {
            return []int{idx, i}
        }
        set[target-n] = i
    }
    return []int{-1,-1}
}
