func canTransform(start string, end string) bool {
    N := len(start)
    i, j := 0, 0
    for i < N && j < N {
        for i < N && start[i] == 'X' {
            i++
        }
        for j < N && end[j] == 'X' {
            j++
        }
        if i >= N && j >= N {
            break
        }
        if (i >= N && j < N) || (i < N) && (j >= N) {
            return false
        }
        if i < N && j < N {
            if start[i] != end[j] || (start[i] == 'L' && i < j) || (start[i] == 'R' && i > j) {
                return false
            }
        }
        i++
        j++
    }
    return true
}
