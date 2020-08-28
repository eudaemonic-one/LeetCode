type MyCircularQueue struct {
    queue []int
    capacity int
    count int
    head int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{make([]int, k), k, 0, 0}
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.count == this.capacity {
        return false
    }
    this.queue[(this.head+this.count)%this.capacity] = value
    this.count++
    return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
    if this.count == 0 {
        return false
    }
    this.head = (this.head + 1) % this.capacity
    this.count--
    return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
    if this.count == 0 {
        return -1
    }
    return this.queue[this.head]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
    if this.count == 0 {
        return -1
    }
    return this.queue[(this.head+this.count-1) % this.capacity]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
    return this.count == 0
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
    return this.count == this.capacity
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
