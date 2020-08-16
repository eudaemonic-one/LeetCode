type FirstUnique struct {
    deque []int
    uniqueMap map[int]bool
}


func Constructor(nums []int) FirstUnique {
    firstUnique := FirstUnique{nums, make(map[int]bool)}
    for _, num := range nums {
        if _, ok := firstUnique.uniqueMap[num]; !ok {
            firstUnique.uniqueMap[num] = true
        } else {
            firstUnique.uniqueMap[num] = false
        }
    }
    return firstUnique
}


func (this *FirstUnique) ShowFirstUnique() int {
    for len(this.deque) > 0 {
        if isUnique, ok := this.uniqueMap[this.deque[0]]; ok && isUnique {
            break
        }
        this.deque = this.deque[1:]
    }
    if len(this.deque) > 0 {
        return this.deque[0]
    }
    return -1
}


func (this *FirstUnique) Add(value int)  {
    if isUnique, ok := this.uniqueMap[value]; !ok {
        this.uniqueMap[value] = true
        this.deque = append(this.deque, value)
    } else if isUnique {
        this.uniqueMap[value] = false
    }
}


/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */
