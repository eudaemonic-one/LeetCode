func threeSumClosest(nums []int, target int) int {
    diff := math.MaxInt64
    
    sort.Ints(nums)
    
    for i := 0; i < len(nums)-2 && diff != 0; i++ {
        j, k := i+1, len(nums)-1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if abs(target-sum) < abs(diff) {
                diff = target - sum
            }
            if sum < target {
                j++
            } else {
                k--
            }
        }
    }
    return target - diff
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
