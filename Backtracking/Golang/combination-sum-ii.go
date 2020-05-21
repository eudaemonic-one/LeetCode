func combinationSum2(candidates []int, target int) [][]int {
    res := make([][]int, 0)
    seq := make([]int, 0)
    sort.Ints(candidates)
    backtrack(candidates, 0, target, &seq, &res)
    return res
}

func backtrack(candidates []int, idx, target int, seq *[]int, res *[][]int) {
    if target == 0 {
        tmp := make([]int, len(*seq))
        copy(tmp, *seq)
        *res = append(*res, tmp)
        return
    }
    
    if target < 0 || idx >= len(candidates) {
        return
    }
    
    for i := idx; i < len(candidates); i++ {
        if i != idx && candidates[i] == candidates[i-1] {
            continue
        }
        
        *seq = append(*seq, candidates[i])
        backtrack(candidates, i+1, target-candidates[i], seq, res)
        (*seq) = (*seq)[:len(*seq)-1]
    }
}
