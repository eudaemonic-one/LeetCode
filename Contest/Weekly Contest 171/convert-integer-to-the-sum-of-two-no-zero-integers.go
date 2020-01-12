func getNoZeroIntegers(n int) []int {
    for i := 1; i <= n/2; i++ {
        a, b := i, n-i
        base := 10
        flag := true
        for a > 0 {
            if a % base == 0 {
                flag = false
                break
            }
            a /= base
        }
        if !flag {
            continue
        }
        for b > 0 {
            if b % base == 0 {
                flag = false
                break
            }
            b /= base
        }
        if !flag {
            continue
        }
        return []int{i, n-i}
    }
    return []int{}
}