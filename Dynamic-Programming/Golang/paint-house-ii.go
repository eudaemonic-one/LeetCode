// Approach #2: DP with tricks

func minCostII(costs [][]int) int {
    if len(costs) == 0 {
        return 0
    }
    n, k := len(costs), len(costs[0])
    
    for i := 1; i < n; i++ {
        firMinIdx, secMinIdx := -1, -1
        for j := 0; j < k; j++ {
            if firMinIdx == -1 || costs[i-1][j] < costs[i-1][firMinIdx] {
                secMinIdx = firMinIdx
                firMinIdx = j
            } else if secMinIdx == -1 || costs[i-1][j] < costs[i-1][secMinIdx] {
                secMinIdx = j
            }
        }
        
        for j := 0; j < k; j++ {
            if j == firMinIdx {
                costs[i][j] += costs[i-1][secMinIdx]
            } else {
                costs[i][j] += costs[i-1][firMinIdx]
            }
        }
    }
    
    res := math.MaxInt64
    for j := 0; j < k; j++ {
        res = min(res, costs[n-1][j])
    }
    return res
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

// Approach #1: Recursion with Memo

// func minCostII(costs [][]int) int {
//     if len(costs) == 0 {
//         return 0
//     }
//     n, k := len(costs), len(costs[0])
//     memo := make(map[string]int)
//     res := math.MaxInt64
//     for i := 0; i < k; i++ {
//         res = min(res, helper(costs, n, k, 0, i, memo))
//     }
//     return res
// }

// func helper(costs [][]int, n, k, idx, color int, memo map[string]int) int {
//     if idx == n-1 {
//         return costs[n-1][color]
//     }
//     key := fmt.Sprintf("%d,%d", idx, color)
//     if val, ok := memo[key]; ok {
//         return val
//     }
//     res := math.MaxInt64
//     for i := 0; i < k; i++ {
//         if i != color {
//             res = min(res, helper(costs, n, k, idx+1, i, memo))
//         }
//     }
//     res += costs[idx][color]
//     memo[key] = res
//     return res
// }

// func min(x, y int) int {
//     if x < y {
//         return x
//     }
//     return y
// }
