func findNumbers(nums []int) int {
    res := 0
    for _, n := range nums {
        digits := 0
        base := 1
        for n / base > 0 {
            base *= 10
            digits++
        }
        if digits % 2 == 0 {
            res++
        }
    }
    return res
}
