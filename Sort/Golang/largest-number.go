type NumStrings []string

func (numStrs NumStrings) Len() int {
    return len(numStrs)
}

func (numStrs NumStrings) Less(i, j int) bool {
    order1 := numStrs[i] + numStrs[j]
    order2 := numStrs[j] + numStrs[i]
    return order1 < order2
}

func (numStrs NumStrings) Swap(i, j int) {
    numStrs[i], numStrs[j] = numStrs[j], numStrs[i]
}

func largestNumber(nums []int) string {
    strs := make([]string, len(nums))
    for i, num := range nums {
        strs[i] = strconv.Itoa(num)
    }
    sort.Sort(NumStrings(strs))
    var sb strings.Builder
    for i := len(strs)-1; i >= 0; i-- {
        sb.WriteString(strs[i])
    }
    res := sb.String()
    offset := 0
    if len(res) > 1 && res[0] == '0' {
        offset = 1
        for i := 1; i < len(res); i++ {
            offset = i
            if res[i] != '0' {
                break
            }
        }
    }
    return res[offset:]
}
