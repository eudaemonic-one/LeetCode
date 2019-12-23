func addBoldTag(s string, dict []string) string {
    wrapIndex := make([]int, 0)
    res := make([]byte, 0)
    for i := 0; i < len(s); i++ {
        maxRange := 0
        for _, substr := range dict {
            if strings.Index(s[i:], substr) == 0 && len(substr) > maxRange {
                maxRange = len(substr)
                start, end := i, i+len(substr)
                if len(wrapIndex) == 0 || start > wrapIndex[len(wrapIndex)-1] {
                    wrapIndex = append(wrapIndex, start)
                    wrapIndex = append(wrapIndex, end)
                }
                if start <= wrapIndex[len(wrapIndex)-1] {
                    if start < wrapIndex[len(wrapIndex)-2] {
                        wrapIndex[len(wrapIndex)-2] = start
                    }
                    if end > wrapIndex[len(wrapIndex)-1] {
                        wrapIndex[len(wrapIndex)-1] = end
                    }
                }
            }
        }
    }
    last := 0
    for i := 0; i < len(wrapIndex); i += 2 {
        start, end := wrapIndex[i], wrapIndex[i+1]
        res = append(res, s[last:start]...)
        res = append(res, "<b>"...)
        res = append(res, s[start:end]...)
        res = append(res, "</b>"...)
        last = end
    }
    if last < len(s) {
        res = append(res, s[last:]...)
    }
    return string(res)
}
