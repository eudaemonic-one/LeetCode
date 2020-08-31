type Item struct {
    value byte
    priority int
    index int
}

type ItemPQ []*Item

func (pq ItemPQ) Len() int {
    return len(pq)
}

func (pq ItemPQ) Less(i, j int) bool {
    return pq[i].priority > pq[j].priority
}

func (pq ItemPQ) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index, pq[j].index = i, j
}

func (pq *ItemPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *ItemPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *ItemPQ) update(item *Item, value byte, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func rearrangeString(s string, k int) string {
    n := len(s)
    counter := make(map[byte]int)
    for i := 0; i < n; i++ {
        counter[s[i]]++
    }
    
    pq := make(ItemPQ, len(counter))
	i := 0
    for key, val := range counter {
		pq[i] = &Item{
			value:    key,
			priority: val,
			index:    i,
		}
		i++
    }
	heap.Init(&pq)

    res := make([]byte, 0)
    queue := make([]*Item, 0)
    
    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item)
        res = append(res, item.value)
        item.priority--
        queue = append(queue, item)
        if len(queue) < k {
            continue
        }
        front := queue[0]
        queue = queue[1:]
        if front.priority > 0 {
            heap.Push(&pq, front)
        }
    }
    
    if len(res) != len(s) {
        return ""
    }
    
    return string(res)
}
