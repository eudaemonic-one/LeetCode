func nextGreaterElement(n int) int {
    a := []byte(strconv.Itoa(n))
    i, j := len(a)-2, len(a)-1
    for i >= 0 && a[i] >= a[i+1] {
        i--
    }
    if i < 0 {
        return -1
    }
    for j >= 0 && a[j] <= a[i] {
        j--
    }
    a[i], a[j] = a[j], a[i]
    for i, j = i+1, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
    if res, err := strconv.Atoi(string(a)); err != nil || res > math.MaxInt32 {
        return -1
    } else {
        return res
    }
}
