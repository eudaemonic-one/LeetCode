func shipWithinDays(weights []int, D int) int {
    l, r := getMax(weights), getSum(weights)+1
    for l < r {
        m := l + (r-l)/2
        if canShip(weights, D, m) {
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

func getSum(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

func canShip(weights []int, D, capacity int) bool {
    i := 0
    for day := 0; day < D; day++ {
        rem := capacity
        for rem - weights[i] >= 0 {
            rem -= weights[i]
            i++
            if i == len(weights) {
                return true
            }
        }
    }
    return false
}
