func jump(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }
    
    maxPos := nums[0]
    maxSteps := nums[0]
    jumps := 1
    for i := 1; i < len(nums); i++ {
        if maxSteps < i {
            jumps++
            maxSteps = maxPos
        }
        maxPos = max(maxPos, i+nums[i])
    }
    return jumps
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
