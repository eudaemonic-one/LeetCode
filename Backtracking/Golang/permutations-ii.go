func permuteUnique(nums []int) [][]int {
    res := make([][]int, 0)
    path := make([]int, 0)
    used := make([]bool, len(nums))
    sort.Ints(nums)
    backtrack(nums, used, path, &res)
    return res
}

func backtrack(nums []int, used []bool, path []int, res *[][]int) {
    if len(path) == len(nums) {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *res = append(*res, tmp)
        return
    }
    for i := 0; i < len(nums); i++ {
        if used[i] == true {
            continue
        }
        if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
            continue
        }
        path = append(path, nums[i])
        used[i] = true
        backtrack(nums, used, path, res)
        path = path[:len(path)-1]
        used[i] = false
    }
}
