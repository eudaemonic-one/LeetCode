func hIndex(citations []int) int {
    N := len(citations)
    count := make([]int, N+1)
    for _, n := range citations {
        count[min(n, N)]++
    }
    n := N
    for sum := count[n]; n > sum; sum += count[n] {
        n--
    }
    return n
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
