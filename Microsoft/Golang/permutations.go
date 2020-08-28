func permute(nums []int) [][]int {
    return backtrack(make([][]int, 0), nums, 0)
}

func backtrack(res [][]int, nums []int, idx int) [][]int {
    if idx == len(nums) {
        tmp := make([]int, len(nums))
        copy(tmp, nums)
        res = append(res, tmp)
        return res
    }
    
    for i := idx; i < len(nums); i++ {
        nums[idx], nums[i] = nums[i], nums[idx]
        res = backtrack(res, nums, idx+1)
        nums[idx], nums[i] = nums[i], nums[idx]
    }
    
    return res
}
