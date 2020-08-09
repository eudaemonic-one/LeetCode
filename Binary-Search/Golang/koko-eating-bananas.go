func minEatingSpeed(piles []int, H int) int {
    l, r := 1, getMax(piles)+1
    for l < r {
        m := l + (r-l)/2
        if canFinish(piles, H, m) {
            r = m
        } else {
            l = m + 1
        }
    }
    return l
}

func getMax(nums []int) int {
    cand := math.MinInt64
    for _, num := range nums {
        if num > cand {
            cand = num
        }
    }
    return cand
}

func canFinish(piles []int, H, K int) bool {
    time := 0
    for _, pile := range piles {
        time += pile / K
        if pile % K > 0 {
            time += 1
        }
    }
    return time <= H
}
