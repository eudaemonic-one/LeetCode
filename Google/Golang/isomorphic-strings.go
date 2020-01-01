func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    dictST := make(map[byte]byte)
    dictTS := make(map[byte]byte)
    for i := 0; i < len(s); i++ {
        if ch, ok := dictST[s[i]]; ok && ch != t[i] {
            return false
        }
        if ch, ok := dictTS[t[i]]; ok && ch != s[i] {
            return false
        }
        dictST[s[i]] = t[i]
        dictTS[t[i]] = s[i]
    }
    return true
}
