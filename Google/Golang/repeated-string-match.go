func repeatedStringMatch(A string, B string) int {
    q := ((len(B)-1)/len(A)) + 1
    for i := 0; i < 2; i++ {
        s := ""
        for j := 0; j < q+i; j++ {
            s += A
        }
        if strings.Contains(s, B) {
            return q+i
        }
    }
    return -1
}
