func threeSum(nums []int) [][]int {
    res := make([][]int, 0)
    n := len(nums)
    if n < 3 {
        return res
    }
    
    sort.Ints(nums)
    
    for i := 0; i < n-2 && nums[i] <= 0; i++ {
        if i > 0 && nums[i-1] == nums[i] {
            continue
        }
        j := i + 1
        k := n - 1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum < 0 {
                j++
            } else if sum > 0 {
                k--
            } else {
                res = append(res, []int{nums[i], nums[j], nums[k]})
                j++
                k--
                for j < k && nums[j-1] == nums[j] {
                    j++
                }
            }
        }
    }
    return res
}
