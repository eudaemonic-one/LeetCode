func subsets(nums []int) [][]int {
    res := make([][]int, 0)
    n := len(nums)
    for k := 0; k <= n; k++ {
        backtrack(nums, 0, k, n, &[]int{}, &res)
    }
    return res
}

func backtrack(nums []int, start, k, n int, path *[]int, res *[][]int) {
    if len(*path) == k {
        tmp := make([]int, k)
        copy(tmp, *path)
        *res = append(*res, tmp)
        return
    }
    for i := start; i < n; i++ {
        *path = append(*path, nums[i])
        backtrack(nums, i+1, k, n, path, res)
        *path = (*path)[:len(*path)-1]
    }
}
