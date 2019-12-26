func kEmptySlots(bulbs []int, K int) int {
    lights := make([]bool, len(bulbs))
    for day, bulb := range bulbs {
        lights[bulb-1] = true
        for i := bulb-2; i >= 0; i-- {
            if lights[i] == true {
                if bulb - i - 2 == K {
                    return day+1
                }
                break
            }
        }
        for i := bulb; i < len(lights); i++ {
            if lights[i] == true {
                if i - bulb == K {
                    return day+1
                }
                break
            }
        }
    }
    return -1
}
