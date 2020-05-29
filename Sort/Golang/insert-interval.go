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

func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {
        return [][]int{newInterval}
    }
    sort.Sort(Intervals(intervals))
    res := [][]int{}
    i, n := 0, len(intervals)
    newStart, newEnd := newInterval[0], newInterval[1]
    
    for i < n && newStart > intervals[i][0] {
        res = append(res, intervals[i])
        i++
    }
    
    if len(res) == 0 || res[len(res)-1][1] < newStart {
        res = append(res, newInterval)
    } else {
        res[len(res)-1][1] = max(res[len(res)-1][1], newEnd)
    }
    
    for i < n {
        itv := intervals[i]
        start, end := itv[0], itv[1]
        if res[len(res)-1][1] < start {
            res = append(res, itv)
        } else {
            res[len(res)-1][1] = max(res[len(res)-1][1], end)
        }
        i++
    }
    
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
