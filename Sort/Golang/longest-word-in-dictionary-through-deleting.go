func findLongestWord(s string, d []string) string {
    sort.Slice(d, func (i, j int) bool {
        if len(d[i]) == len(d[j]) {
            return d[i] < d[j]
        }
        return len(d[i]) > len(d[j])
    })
    for _, word := range d {
        if isSubsequence(word, s) {
            return word
        }
    }
    return ""
}

func isSubsequence(s1, s2 string) bool {
    i, j := 0, 0
    for i < len(s1) && j < len(s2) {
        if s1[i] == s2[j] {
            i++
        }
        j++
    }
    return i == len(s1)
}
