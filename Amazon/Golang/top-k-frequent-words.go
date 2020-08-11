// Approach 2: Heap

type Pair struct {
    key string
    val int
}

type PairHeap []Pair

func (h PairHeap) Len() int {
    return len(h)
}

func (h PairHeap) Less(i, j int) bool {
    if h[i].val == h[j].val {
        return h[i].key > h[j].key
    }
    return h[i].val < h[j].val
}

func (h PairHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x interface{}) {
    *h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func topKFrequent(words []string, k int) []string {
    counter := make(map[string]*Pair)
    for _, word := range words {
        if pair, ok := counter[word]; ok {
            pair.val++
        } else {
            counter[word] = &Pair{word, 1}
        }
    }
    
    h := &PairHeap{}
	heap.Init(h)
    for _, pair := range counter {
        heap.Push(h, *pair)
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    
    res := make([]string, 0)
    for i := 0; i < k && h.Len() > 0; i++ {
        res = append(res, heap.Pop(h).(Pair).key)
    }
    for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
        res[l], res[r] = res[r], res[l]
    }
    return res
}

// Approach 1: Sorting

// func topKFrequent(words []string, k int) []string {
//     counter := make(map[string]*Pair)
//     for _, word := range words {
//         if pair, ok := counter[word]; ok {
//             pair.val++
//         } else {
//             counter[word] = &Pair{word, 1}
//         }
//     }
//     pairs := make([]Pair, 0)
//     for _, pair := range counter {
//         pairs = append(pairs, *pair)
//     }
//     sort.SliceStable(pairs, func (i, j int) bool {
//         if pairs[i].val == pairs[j].val {
//             return pairs[i].key < pairs[j].key
//         }
//         return pairs[i].val > pairs[j].val
//     })
//     res := make([]string, 0)
//     for i := 0; i < k && i < len(pairs); i++ {
//         res = append(res, pairs[i].key)
//     }
//     return res
// }
