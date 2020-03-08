func generateTheString(n int) string {
    res := make([]byte, 0)
    if n % 2 == 0 {
        for i := 0; i < n-1; i++ {
            res = append(res, 'a')
        }
        res = append(res, 'b')
    } else {
        for i := 0; i < n; i++ {
            res = append(res, 'a')
        }
    }
    return string(res)
}