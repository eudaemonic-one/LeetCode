type Intervals [][]int

func (itvs Intervals) Len() int {
    return len(itvs)
}

func (itvs Intervals) Less(i, j int) bool {
    return itvs[i][0] < itvs[j][0]
}

func (itvs Intervals) Swap(i, j int) {
    itvs[i], itvs[j] = itvs[j], itvs[i]
}

func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return intervals
    }
    sort.Sort(Intervals(intervals))
    res := [][]int{[]int{intervals[0][0], intervals[0][1]}}
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] <= res[len(res)-1][1] {
            res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
        } else {
            res = append(res, intervals[i])
        }
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
