type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func connectSticks(sticks []int) int {
    if len(sticks) == 0 {
        return 0
    }
    
    h := &IntHeap{} 
    heap.Init(h)
    for _, stick := range sticks {
        heap.Push(h, stick)
    }
    
    res := 0
    
    for h.Len() >= 2 {
        x := heap.Pop(h).(int)
        y := heap.Pop(h).(int)
        heap.Push(h, x+y)
        res += x + y
    }
    
    return res
}
