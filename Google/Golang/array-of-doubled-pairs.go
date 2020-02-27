func canReorderDoubled(A []int) bool {
    if len(A) % 2 == 1 {
        return false
    }
    counter := make(map[int]int)
    for _, x := range A {
        counter[x] += 1
    }
    sort.Slice(A, func(i, j int) bool {
        if (A[i] < 0 && A[j] < 0) {
            return A[i] > A[j]
        } else {
            return A[i] < A[j]
        }
    })
    for _, x := range A {
        if counter[x] == 0 {
            continue
        }
        if counter[2*x] == 0 {
            return false
        }
        counter[x] -= 1
        counter[2*x] -= 1
    }
    return true
}
