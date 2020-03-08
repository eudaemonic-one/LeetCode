func numTimesAllBlue(light []int) int {
    const OFF = 0
    const ON = 1
    const BLUE = 2
    res := 0
    lights := len(light)
    states := make([]int, lights)
    for i := 0; i < lights; i++ {
        states[light[i]-1] = ON
        for j := 0; j < lights; j++ {
            if states[j] == OFF {
                break
            } else if states[j] == ON {
                states[j] = BLUE
            }
        }
        flag := true
        for k := 0; k < lights; k++ {
            if states[k] == ON {
                flag = false
                break
            }
        }
        if flag {
            res += 1
        }
    }
    return res
}