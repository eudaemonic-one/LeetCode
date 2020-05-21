func subsetsWithDup(nums []int) [][]int {
    res := make([][]int, 0)
    sort.Ints(nums)
    backtrack(nums, 0, len(nums), &[]int{}, &res)
    return res
}

func backtrack(nums []int, start, n int, path *[]int, res *[][]int) {
    if len(*path) <= n {
        tmp := make([]int, len(*path))
        copy(tmp, *path)
        *res = append(*res, tmp)
    }
    for i := start; i < n; i++ {
        if i > start && nums[i] == nums[i-1] {
            continue
        }
        *path = append(*path, nums[i])
        backtrack(nums, i+1, n, path, res)
        *path = (*path)[:len(*path)-1]
    }
}
