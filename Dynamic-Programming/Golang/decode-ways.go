func numDecodings(s string) int {
    if len(s) == 0 {
        return 0
    }
    a, b, c := 1, 1, 1
    if s[0] == '0' {
        if len(s) == 1 {
            return 0
        } else {
            b = 0
        }
    }
    
    for i := 2; i <= len(s); i++ {
        c = 0
        if s[i-1] != '0' {
            c += b
        }
        twoDigit, _ := strconv.Atoi(s[i-2:i])
        if 10 <= twoDigit && twoDigit <= 26 {
            c += a
        }
        a = b
        b = c
    }
    
    return c
}
