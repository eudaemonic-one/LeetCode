type Intervals [][2]int

func (intervals Intervals) Len() int {
    return len(intervals)
}

func (intervals Intervals) Less(i, j int) bool {
    if intervals[i][0] == intervals[j][0] {
        return intervals[i][1] < intervals[j][1]
    } else {
        return intervals[i][0] < intervals[j][0]
    }
}

func (intervals Intervals) Swap(i, j int) {
    intervals[i], intervals[j] = intervals[j], intervals[i]
}

func addBoldTag(s string, dict []string) string {
    intervals := Intervals{}
    res := make([]byte, 0)
    for i := 0; i < len(s); i++ {
        for _, substr := range dict {
            if strings.Index(s[i:], substr) == 0 {
                intervals = append(intervals, [2]int{i,i+len(substr)})
            }
        }
    }
    sort.Sort(intervals)
    merged := Intervals{}
    for _, interval := range intervals {
        if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
            merged = append(merged, interval)
        } else {
            merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
        }
    }
    last := 0
    for _, interval := range merged {
        start, end := interval[0], interval[1]
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

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
