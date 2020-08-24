// Approach 2: Chronological Ordering

func minMeetingRooms(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }
    
    startTimes := make([]int, len(intervals))
    endTimes := make([]int, len(intervals))
    for i := 0; i < len(intervals); i++ {
        startTimes[i] = intervals[i][0]
        endTimes[i] = intervals[i][1]
    }
    
    sort.Ints(startTimes)
    sort.Ints(endTimes)
    
    usedRooms := 0
    start, end := 0, 0
    
    for start < len(intervals) {
        if startTimes[start] >= endTimes[end] {
            usedRooms--
            end++
        }
        usedRooms++
        start++
    }
    
    return usedRooms
}

// Approach 1: Priority Queues

// type Intervals [][]int

// func (itvls Intervals) Len() int { return len(itvls) }
// func (itvls Intervals) Less(i, j int) bool { return itvls[i][0] < itvls[j][0] }
// func (itvls Intervals) Swap(i, j int){ itvls[i], itvls[j] = itvls[j], itvls[i] }

// type IntHeap []int

// func (h IntHeap) Len() int { return len(h) }
// func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h IntHeap) Swap(i, j int){ h[i], h[j] = h[j], h[i] }

// func (h *IntHeap) Push(x interface{}) {
//     *h = append(*h, x.(int))
// }

// func (h *IntHeap) Pop() interface{} {
//     old := *h
//     n := len(old)
//     x := old[n-1]
//     old = old[0:n-1]
//     *h = old
//     return x
// }

// func minMeetingRooms(intervals [][]int) int {    
//     if len(intervals) == 0 {
//         return 0
//     }
//     sort.Sort(Intervals(intervals))
    
//     rooms := IntHeap{intervals[0][1]}
//     heap.Init(&rooms)
//     for i := 1; i < len(intervals); i++ {
//         if intervals[i][0] >= rooms[0] {
//             heap.Pop(&rooms)
//         }
//         heap.Push(&rooms, intervals[i][1])
//     }
//     return len(rooms)
// }
