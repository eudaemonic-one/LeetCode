type Interval struct {
    start int
    end int
}

type MyCalendarTwo struct {
    single []Interval
    double []Interval
}


func Constructor() MyCalendarTwo {
    return MyCalendarTwo{make([]Interval,0),make([]Interval,0)}
}


func (this *MyCalendarTwo) Book(start int, end int) bool {
    for _, itv := range this.double {
        if end > itv.start && start < itv.end {
            return false
        }
    }
    for _, itv := range this.single {
        if end > itv.start && start < itv.end {
            this.double = append(this.double, Interval{max(start, itv.start),min(end,itv.end)})
        }
    }
    this.single = append(this.single, Interval{start, end})
    return true
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}


/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
