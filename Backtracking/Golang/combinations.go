func combine(n int, k int) [][]int {
    res := make([][]int, 0)
    backtrack(n, k, 1, &[]int{}, &res)
    return res
}

func backtrack(n, k, start int, comb *[]int, res *[][]int) {
    if len(*comb) == k {
        tmp := make([]int, k)
        copy(tmp, *comb)
        *res = append(*res, tmp)
        return
    }
    for i := start; i <= n; i++ {
        *comb = append(*comb, i)
        backtrack(n, k, i+1, comb, res)
        *comb = (*comb)[:len(*comb)-1]
    }
}
