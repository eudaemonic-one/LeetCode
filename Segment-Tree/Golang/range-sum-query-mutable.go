type NumArray struct {
    arr []int
    st []int
    n   int
}

func Constructor(nums []int) NumArray {
    n := len(nums)
    if n == 0 {
        return NumArray{nums, make([]int, 0), 0}
    }
    x := math.Ceil(math.Log2(float64(n)))
    maxSize := 2*int(math.Pow(2, x)) - 1
    st := make([]int, maxSize)
    constructTreeUtil(nums, 0, n-1, &st, 0)
    return NumArray{nums, st, n}
}

func constructTreeUtil(arr []int, ss, se int, st *[]int, si int) int {
    if ss == se {
        (*st)[si] = arr[ss]
        return arr[ss]
    }
    mid := ss + (se-ss)/2
    (*st)[si] = constructTreeUtil(arr, ss, mid, st, si*2+1) + constructTreeUtil(arr, mid+1, se, st, si*2+2)
    return (*st)[si]
}

func (this *NumArray) Update(i int, val int)  {
    if i < 0 || i >= this.n {
        return
    }
    diff := val - this.arr[i]
    this.arr[i] = val
    updateUtil(&this.st, 0, this.n-1, i, diff, 0)
}

func updateUtil(st *[]int, ss, se, i, diff, si int) {
    if i < ss || i > se {
        return
    }
    (*st)[si] += diff
    if ss != se {
        mid := ss + (se-ss)/2
        updateUtil(st, ss, mid, i, diff, 2*si+1)
        updateUtil(st, mid+1, se, i, diff, 2*si+2)
    }
}

func (this *NumArray) SumRange(i int, j int) int {
    if i < 0 || j >= this.n || i > j {
        return -1
    }
    return getSumUtil(this.st, 0, this.n-1, i, j, 0)
}

func getSumUtil(st []int, ss, se, qs, qe, si int) int {
    if qs <= ss && qe >= se {
        return st[si]
    }
    if qs > se || qe < ss {
        return 0
    }
    mid := ss + (se-ss)/2
    return getSumUtil(st, ss, mid, qs, qe, 2*si+1) + getSumUtil(st, mid+1, se, qs, qe, 2*si+2)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
