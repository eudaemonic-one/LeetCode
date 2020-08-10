func twoSum(nums []int, target int) []int {
    table := make(map[int]int)
    for i, n := range nums {
        if idx, ok := table[n]; ok {
            return []int{idx, i}
        }
        table[target-n] = i
    }
    return []int{-1,-1}
}
