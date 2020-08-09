type Intervals [][]int

func (itvs Intervals) Len() int {
    return len(itvs)
}

func (itvs Intervals) Less(i, j int) bool {
    return itvs[i][1] < itvs[j][1]
}

func (itvs Intervals) Swap(i, j int) {
    itvs[i], itvs[j] = itvs[j], itvs[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }
    sort.Sort(Intervals(intervals))
    overlapped := 1
    lastEnd := intervals[0][1]
    for _, itv := range intervals {
        start, end := itv[0], itv[1]
        if start >= lastEnd {
            overlapped++
            lastEnd = end
        }
    }
    return len(intervals) - overlapped
}
