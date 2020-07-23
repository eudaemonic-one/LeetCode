func totalFruit(tree []int) int {
    res := 0
    counter := make(map[int]int)
    l, r := 0, 0
    for r < len(tree) {
        counter[tree[r]]++
        r++
        for len(counter) > 2 {
            counter[tree[l]]--
            if counter[tree[l]] == 0 {
                delete(counter, tree[l])
            }
            l++
        }
        if r-l > res {
            res = r - l
        }
    }
    return res
}
