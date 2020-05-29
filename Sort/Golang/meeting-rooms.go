type Intervals [][]int

func (itvs Intervals) Len() int { return len(itvs) }
func (itvs Intervals) Swap(i, j int){ itvs[i], itvs[j] = itvs[j], itvs[i] }
func (itvs Intervals) Less(i, j int) bool { return itvs[i][0] < itvs[j][0] }

func canAttendMeetings(intervals [][]int) bool {
    sort.Sort(Intervals(intervals))
    for i := 0; i < len(intervals)-1; i++ {
        if intervals[i][1] > intervals[i+1][0] {
            return false
        }
    }
    return true
}
