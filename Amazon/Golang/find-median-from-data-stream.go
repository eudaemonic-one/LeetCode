type MedianFinder struct {
    left  *IntHeap
    right *IntHeap
}

type IntHeap []int // Min-Heap

func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Top() interface{} {
    if len(*h) == 0 {
        return nil
    }
    return (*h)[0]
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{&IntHeap{}, &IntHeap{}}
}


func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.left, num)
    heap.Push(this.right, -heap.Pop(this.left).(int))
    if this.left.Len() < this.right.Len() {
        heap.Push(this.left, -heap.Pop(this.right).(int))
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.left.Len() > this.right.Len() {
        return float64(this.left.Top().(int))
    }
    return float64(this.left.Top().(int)) + (-float64(this.right.Top().(int)) - float64(this.left.Top().(int))) / 2
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
