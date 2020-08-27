// Approach 2: Dynamic Programming

func numTimesAllBlue(light []int) int {
    res := 0
    right := 0
    for i := 0; i < len(light); i++ {
        right = max(right, light[i])
        if right == i+1 {
            res++
        }
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

// Approach 1: Brute Force

// const OFF = 0
// const ON = 1
// const BLUE = 2

// func numTimesAllBlue(light []int) int {
//     res := 0
//     n := len(light)
//     states := make([]int, n)
//     for i := 0; i < n; i++ {
//         states[light[i]-1] = ON
//         for j := 0; j < n; j++ {
//             if states[j] == OFF {
//                 break
//             } else if states[j] == ON {
//                 states[j] = BLUE
//             }
//         }
//         flag := true
//         for k := 0; k < n; k++ {
//             if states[k] == ON {
//                 flag = false
//                 break
//             }
//         }
//         if flag {
//             res++
//         }
//     }
//     return res
// }
