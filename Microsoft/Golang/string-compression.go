func compress(chars []byte) int {
    if len(chars) <= 1 {
        return len(chars)
    }
    anchor, write := 0, 0
    for read := 0; read < len(chars); read++ {
        if read+1 == len(chars) || chars[read] != chars[read+1] {
            chars[write] = chars[anchor]
            write++
            if read > anchor {
                digits := strconv.Itoa(read-anchor+1)
                for i := 0; i < len(digits); i++ {
                    chars[write] = digits[i]
                    write++
                }
            }
            anchor = read + 1
        }
    }
    return write
}
