func permute(nums []int) [][]int {
    res := make([][]int, 0)
    backtrack(nums, 0, &res)
    return res
}

func backtrack(nums []int, idx int, res *[][]int) {
    if idx == len(nums) {
        tmp := make([]int, len(nums))
        copy(tmp, nums)
        *res = append(*res, tmp)
    }
    for i := idx; i < len(nums); i++ {
        nums[idx], nums[i] = nums[i], nums[idx]
        backtrack(nums, idx+1, res)
        nums[idx], nums[i] = nums[i], nums[idx]
    }
}
