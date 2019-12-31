type LRUCache struct {
    mapCache map[int](*list.Element)
    listCache *list.List
    size int
    capacity int
}


type LRUElement struct {
    key int
    value int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{make(map[int](*list.Element)), list.New(), 0, capacity}
}


func (this *LRUCache) Get(key int) int {
    if val, ok := this.mapCache[key]; ok {
        this.listCache.MoveToFront(val)
        return val.Value.(*LRUElement).value
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if val, ok := this.mapCache[key]; ok {
        val.Value.(*LRUElement).value = value
        this.listCache.MoveToFront(val)
        return
    }
    
    if len(this.mapCache) >= this.capacity {
        elem := this.listCache.Back()
        this.listCache.Remove(elem)
        delete(this.mapCache, elem.Value.(*LRUElement).key)
	}

	item := &LRUElement{key: key, value: value}
	element := this.listCache.PushFront(item)
	this.mapCache[key] = element
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
