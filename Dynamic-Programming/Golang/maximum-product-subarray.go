func maxProduct(nums []int) int {
    res := nums[0]
    curMax, curMin := nums[0], nums[0]
    
    for i := 1; i < len(nums); i++ {
        if nums[i] < 0 {
            curMax, curMin = curMin, curMax
        }
        
        if nums[i] > curMax * nums[i] {
            curMax = nums[i]
        } else {
            curMax = curMax * nums[i]
        }
        
        if nums[i] < curMin * nums[i] {
            curMin = nums[i]
        } else {
            curMin = curMin * nums[i]
        }
        
        if curMax > res {
            res = curMax
        }
    }
    
    return res
}
