func repeatedSubstringPattern(s string) bool {
    if len(s) <= 1 {
        return false
    }
    return strings.Index((s+s)[1:len(s)*2-1], s) != -1
}
