func repeatedSubstringPattern(s string) bool {
    for i := 0; i < len(s)-1; i++ {
        if len(s) % (i+1) != 0 {
            continue
        }
        substr, t := s[:i+1], ""
        for j := 0; j < (len(s)/(i+1)); j++ {
            t += substr
        }
        if s == t {
            return true
        }
    }
    return false
}
