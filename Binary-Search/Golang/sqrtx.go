func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    l, r := 2, x/2
    for l <= r {
        m := l + (r-l)/2
        num := m * m
        if num > x {
            r = m - 1
        } else if num < x {
            l = m + 1
        } else {
            return m
        }
    }
    return r
}
