func maximumGap(nums []int) int {
    if len(nums) < 2 {
        return 0
    }
    max := 0
    for _, num := range nums {
        if num > max {
            max = num
        }
    }
    
    n := len(nums)
    for exp := 1; max/exp > 0; exp *= 10 {
        countSort(nums, n, exp)
    }
    
    res := 0
    for i := 1; i < n; i++ {
        if diff := nums[i] - nums[i-1]; diff > res {
            res = diff
        }
    }
    return res
}

func countSort(nums []int, n int, exp int) {
    res := make([]int, n)
    count := make([]int, 10)
    
    for i := 0; i < n; i++ {
        count[(nums[i]/exp)%10]++
    }
    
    for i := 1; i < 10; i++ {
        count[i] += count[i-1]
    }
    
    for i := n-1; i >= 0; i-- {
        res[count[(nums[i]/exp)%10]-1] = nums[i]
        count[(nums[i]/exp)%10]--
    }
    
    for i := 0; i < n; i++ {
        nums[i] = res[i]
    }
}
