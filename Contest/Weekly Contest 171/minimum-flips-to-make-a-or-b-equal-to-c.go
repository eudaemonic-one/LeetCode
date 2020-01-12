func minFlips(a int, b int, c int) int {
    flips := 0
    for a > 0 || b > 0 || c > 0 {
        bitA := a & 0x01
        bitB := b & 0x01
        bitC := c & 0x01
        if bitC == 0x01 {
            if bitA == 0x01 || bitB == 0x01 {
                a = a >> 1
                b = b >> 1
                c = c >> 1
                continue
            }
            flips++
        } else {
            if bitA == 0x01 {
                flips++
            }
            if bitB == 0x01 {
                flips++
            }
        }
        a = a >> 1
        b = b >> 1
        c = c >> 1
    }
    return flips
}