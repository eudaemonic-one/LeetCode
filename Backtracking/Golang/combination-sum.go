func combinationSum(candidates []int, target int) [][]int {
    res := make([][]int, 0)
    seq := make([]int, 0)
    backtrack(candidates, 0, target, &seq, &res)
    return res
}

func backtrack(candidates []int, idx, target int, seq *[]int, res *[][]int) {
    if target < 0 {
        return
    }
    
    if target == 0 {
        tmp := make([]int, len(*seq))
        copy(tmp, *seq)
        *res = append(*res, tmp)
        return
    }
    
    for i := idx; i < len(candidates); i++ {
        *seq = append(*seq, candidates[i])
        backtrack(candidates, i, target-candidates[i], seq, res)
        *seq = (*seq)[:len(*seq)-1]
    }
}
