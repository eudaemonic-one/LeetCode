func combinationSum(candidates []int, target int) [][]int {
    res := make([][]int, 0)
    seq := make([]int, 0)
    backtrack(candidates, 0, target, &seq, &res)
    return res
}

func backtrack(candidates []int, idx, target int, seq *[]int, res *[][]int) {
    if idx >= len(candidates) { // Hit end state
        if target == 0 {
            ans := make([]int, len(*seq))
            copy(ans, *seq)
            *res = append(*res, ans)
        }
        return
    } else if target < 0 { // Prun if target less than zero
        return
    }
    
    // If not append current candidate at index idx
    backtrack(candidates, idx+1, target, seq, res)
    
    // If append current candidate at index idx (unlimited number of times)
    i := 0
    for target >= 0 {
        target -= candidates[idx]
        *seq = append(*seq, candidates[idx])
        i++
        backtrack(candidates, idx+1, target, seq, res)
    }
    for ;i > 0; i-- {
        (*seq) = (*seq)[:len(*seq)-1]
    }
}
