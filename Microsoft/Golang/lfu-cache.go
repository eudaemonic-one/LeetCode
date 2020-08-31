type DoublyLinkedList struct {
    dummy *Node
    size int
}


func NewDoublyLinkedList() *DoublyLinkedList {
    l := &DoublyLinkedList{}
    l.dummy = &Node{0, 0, 0, nil, nil}
    l.dummy.prev = l.dummy
    l.dummy.next = l.dummy
    return l
}


func (this *DoublyLinkedList) Append(node *Node) {
    if node == nil {
        return
    }
    node.next = this.dummy.next
    node.prev = this.dummy
    node.next.prev = node
    this.dummy.next = node
    this.size++
}


func (this *DoublyLinkedList) Pop(node *Node) *Node {
    if this.size == 0 {
        return nil
    }
    if node == nil {
        node = this.dummy.prev
    }
    node.prev.next = node.next
    node.next.prev = node.prev
    this.size--
    return node
}


type Node struct {
    key int
    val int
    frequency int
    prev, next *Node
}


type LFUCache struct {
    nodeTable map[int]*Node
    frequencyTable map[int]*DoublyLinkedList
    minFrequency int
    size int
    capacity int
}


func Constructor(capacity int) LFUCache {
    return LFUCache{
        make(map[int]*Node),
        make(map[int]*DoublyLinkedList),
        0, 0, capacity}
}


func (this *LFUCache) Get(key int) int {
    if node, ok := this.nodeTable[key]; !ok || node == nil {
        return -1
    } else {
        this.update(node)
        return node.val
    }
}


func (this *LFUCache) update(node *Node) {
    frequency := node.frequency
    this.frequencyTable[frequency].Pop(node)
    if list, _ := this.frequencyTable[frequency]; list.size == 0 && this.minFrequency == frequency {
        this.minFrequency += 1
    }
    node.frequency++
    if this.frequencyTable[node.frequency] == nil {
        this.frequencyTable[node.frequency] = NewDoublyLinkedList()
    }
    this.frequencyTable[node.frequency].Append(node)
}


func (this *LFUCache) Put(key int, value int)  {
    if this.capacity == 0 {
        return
    }
    if node, ok := this.nodeTable[key]; ok {
        node.val = value
        this.update(node)
    } else {
        if this.size == this.capacity {
            oldNode := this.frequencyTable[this.minFrequency].Pop(nil)
            delete(this.nodeTable, oldNode.key)
            this.size--
        }
        
        newNode := &Node{key, value, 1, nil, nil}
        this.nodeTable[key] = newNode
        if this.frequencyTable[1] == nil {
            this.frequencyTable[1] = NewDoublyLinkedList()
        }
        this.frequencyTable[1].Append(newNode)
        this.minFrequency = 1
        this.size++
    }
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
