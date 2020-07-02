func findTargetSumWays(nums []int, S int) int {
    res := 0
    backtrack(nums, 0, S, &res)
    return res
}

func backtrack(nums []int, i, target int, res *int) {
    if i == len(nums) {
        if target == 0 {
            *res++
        }
        return
    }
    backtrack(nums, i+1, target-nums[i], res)
    backtrack(nums, i+1, target+nums[i], res)
}
