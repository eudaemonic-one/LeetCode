func reorderLogFiles(logs []string) []string {
    sort.SliceStable(logs, func(i int, j int) bool {
        words1 := strings.SplitN(logs[i], " ", 2)
        words2 := strings.SplitN(logs[j], " ", 2)
        isDigit1 := '0' <= words1[1][0] && words1[1][0] <= '9'
        isDigit2 := '0' <= words2[1][0] && words2[1][0] <= '9'
        if !isDigit1 && !isDigit2 {
            if words1[1] == words2[1] {
                return words1[0] < words2[0]
            }
            return words1[1] < words2[1]
        } else if isDigit1 && isDigit2 {
            return false
        } else {
            return isDigit2
        }
    })
    return logs
}
