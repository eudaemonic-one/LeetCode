func getFactors(n int) [][]int {
    res := make([][]int, 0)
    backtrack(n, &[]int{}, &res)
    return res
}

func backtrack(n int, path *[]int, res *[][]int) {
    factors := len(*path)
    if n == 1 {
        if factors > 1 {
            tmp := make([]int, len(*path))
            copy(tmp, *path)
            *res = append(*res, tmp)
        }
        return
    }
    
    start := 2
    if factors > 0 {
        start = (*path)[factors-1]
    }
    
    for i := start; i <= n; i++ {
        if n % i == 0 {
            *path = append(*path, i)
            backtrack(n/i, path, res)
            *path = (*path)[:factors]
        }
    }
}
