func canTransform(start string, end string) bool {
    reducedStart, reducedEnd := "", ""
    for i := 0; i < len(start); i++ {
        if start[i] != 'X' {
            reducedStart += string(start[i])
        }
        if end[i] != 'X' {
            reducedEnd += string(end[i])
        }
    }
    if reducedStart != reducedEnd {
        return false
    }
    LinEnd := make([]int, 0)
    RinEnd := make([]int, 0)
    for i := 0; i < len(end); i++ {
        if end[i] == 'L' {
            LinEnd = append(LinEnd, i)
        } else if end[i] == 'R' {
            RinEnd = append(RinEnd, i)
        }
    }
    l, r := 0, 0
    for i := 0; i < len(start); i++ {
        if start[i] == 'L' {
            if i < LinEnd[l] {
                return false
            }
            l++
        } else if start[i] == 'R' {
            if i > RinEnd[r] {
                return false
            }
            r++
        }
    }
    return true
}
