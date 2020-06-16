func myPow(x float64, n int) float64 {
    res := 1.0
    if n < 0 {
        x = 1.0 / x
        n = -n
    }
    for n > 0 {
        if n & 1 == 1 {
            res *= x
        }
        x = x * x
        n /= 2
    }
    return res
}
