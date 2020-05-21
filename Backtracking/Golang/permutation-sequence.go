func getPermutation(n int, k int) string {
    factorials := make([]int, n)
    nums := []int{1}
    
    factorials[0] = 1
    for i := 1; i < n; i++ {
        // generate factorial system bases 0!, 1!, ..., (n-1)!
        factorials[i] = factorials[i-1] * i
        // generate nums 1, 2, ..., n
        nums = append(nums, i+1)
    }
    
    // fit k in the interval 0 ... (n!-1)
    k--
    
    // compute factorial representation of k
    res := ""
    for i := n - 1; i >= 0; i-- {
        idx := k / factorials[i]
        k -= idx * factorials[i]
        res += strconv.Itoa(nums[idx])
        for j := idx; j < len(nums)-1; j++ {
            nums[j] = nums[j+1]
        }
        nums = nums[:len(nums)-1]
    }
    
    return res
}
