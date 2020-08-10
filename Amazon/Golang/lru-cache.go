type LRUCache struct {
    mapCache  map[int](*list.Element)
    listCache *list.List
    capacity  int
}

type LRUElement struct {
    Key int
    Val int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{make(map[int](*list.Element)), list.New(), capacity}
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.mapCache[key]; ok {
        this.listCache.MoveToFront(node)
        return node.Value.(*LRUElement).Val
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.mapCache[key]; ok {
        node.Value.(*LRUElement).Val = value
        this.listCache.MoveToFront(node)
    } else {
        if len(this.mapCache) >= this.capacity {
            last := this.listCache.Back()
            this.listCache.Remove(last)
            delete(this.mapCache, last.Value.(*LRUElement).Key)
        }
        elem := &LRUElement{key, value}
        newNode := this.listCache.PushFront(elem)
        this.mapCache[key] = newNode
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
