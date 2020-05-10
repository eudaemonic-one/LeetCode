func mySqrt(x int) int {
    if x == 0 {
        return 0
    }
    l, r := 0, x
    for l < r {
        m :=  l + int(math.Round(float64(r - l) / 2))
        if m <= x / m {
            l = m
        } else {
            r = m - 1
        }
    }
    return l
}

