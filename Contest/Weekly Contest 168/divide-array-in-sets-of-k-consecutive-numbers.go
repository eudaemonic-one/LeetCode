func isPossibleDivide(nums []int, k int) bool {
    if len(nums) % k != 0 {
        return false
    }
    max := 0
    for _, n := range nums {
        if n > max {
            max = n
        }
    }
    counter := make([]int, max+1)
    for _, n := range nums {
        counter[n]++
    }
    n := len(nums) / k
    l, r := 1, 1
    for i := 0; i < n && r < len(counter); {
        if counter[l] > 0 {
            prev := 0
            r = l
            for j := 0; j < k; j++ {
                if counter[r] > 1 {
                    prev = r
                }
                if counter[r] == 0 {
                    return false
                }
                counter[r]--
                r++
            }
            i++
            r = prev
        } else {
            l++
        }
    }
    for i := 0; i < len(counter); i++ {
        if counter[i] > 0 {
            return false
        }
    }
    return true
}
