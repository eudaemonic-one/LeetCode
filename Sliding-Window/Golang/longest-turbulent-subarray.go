func maxTurbulenceSize(A []int) int {
    res := 1
    l, r := 0, 1
    for ; r < len(A); r++ {
        c := cmp(A[r-1], A[r])
        
        if c == 0 {
            l = r
        } else if r == len(A)-1 || c * cmp(A[r], A[r+1]) != -1 {
            if r - l + 1 > res {
                res = r - l + 1
            }
            l = r
        }
    }
    return res
}

func cmp(x, y int) int {
    if x < y {
        return -1
    } else if x > y {
        return 1
    } else {
        return 0
    }
}
