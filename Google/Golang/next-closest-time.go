func nextClosestTime(time string) string {
    res := ""
    choice := make(map[int]bool)
    for i := 0; i < len(time); i++ {
        if time[i] != ':' {
            digit, _ := strconv.Atoi(string(time[i]))
            choice[digit] = true
        }
    }
    h, _ := strconv.Atoi(string(time[0:2]))
    m, _ := strconv.Atoi(string(time[3:5]))
    cur := h * 60 + m
    for true {
        cur = (cur + 1) % (24 * 60)
        // digit 4
        if _, ok := choice[cur%60%10]; !ok {
            continue
        }
        // digit 3
        if _, ok := choice[(cur%60)/10]; !ok {
            continue
        }
        // digit 2
        if _, ok := choice[(cur/60)%10]; !ok {
            continue
        }
        // digit 1
        if _, ok := choice[(cur/60)/10]; !ok {
            continue
        }
        hour := strconv.Itoa(cur/60)
        min := strconv.Itoa(cur%60)
        if len(hour) == 1 {
            res += "0"
        }
        res += hour + ":"
        if len(min) == 1 {
            res += "0"
        }
        res += min
        return res
    }
    return ""
}
