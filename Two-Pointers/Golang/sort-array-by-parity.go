func sortArrayByParity(A []int) []int {
    slow := 0
    for slow < len(A) && A[slow] % 2 != 1 {
        slow++
    }
    fast := slow+1
    for fast < len(A) {
        if A[fast] % 2 == 0 {
            A[slow], A[fast] = A[fast], A[slow]
            slow++
        }
        fast++
    }
    return A
}
