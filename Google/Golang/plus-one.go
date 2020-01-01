func plusOne(digits []int) []int {
    sum := digits[len(digits)-1] + 1
    carry, remain := sum/10, sum%10
    res := []int{remain}
    for i := len(digits)-2; i >= 0; i-- {
        sum := digits[i] + carry
        carry, remain = sum/10, sum%10
        res = append([]int{remain}, res...)
    }
    if carry > 0 {
        res = append([]int{carry}, res...)
    }
    return res
}
