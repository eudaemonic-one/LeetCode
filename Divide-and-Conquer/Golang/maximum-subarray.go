func maxSubArray(nums []int) int {
    return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) int {
    if (left == right) {
        return nums[left]
    }
    pi := left + (right-left)/2
    leftSum := helper(nums, left, pi)
    rightSum := helper(nums, pi+1, right)
    crossSum := crossSum(nums, left, right, pi)
    return max(max(leftSum, rightSum), crossSum)
}

func crossSum(nums []int, left, right, pi int) int {
    if (left == right) {
        return nums[left]
    }
    leftSubSum := math.MinInt64
    currSum := 0
    for i := pi; i >= left; i-- {
        currSum += nums[i]
        leftSubSum = max(leftSubSum, currSum)
    }
    rightSubSum := math.MinInt64
    currSum = 0
    for i := pi+1; i <= right; i++ {
        currSum += nums[i]
        rightSubSum = max(rightSubSum, currSum)
    }
    return leftSubSum + rightSubSum
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
