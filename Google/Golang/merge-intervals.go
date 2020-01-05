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
    if len(intervals) <= 1 {
        return intervals
    }
    sort.Sort(Intervals(intervals))
    merged := make([][]int, 0)
    curr, next := intervals[0], []int{}
    for i := 1; i < len(intervals); i++ {
        next = intervals[i]
        if next[0] <= curr[1] {
            curr[1] = max(curr[1],next[1])
        } else {
            merged = append(merged, curr)
            curr = next
        }
        if i == len(intervals)-1 {
            merged = append(merged, curr)
        }
    }
    return merged
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
