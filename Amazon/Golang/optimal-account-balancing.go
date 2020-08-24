func minTransfers(transactions [][]int) int {
    balanceTable := make(map[int]int)  
    for _, trx := range transactions {
        x, y, z := trx[0], trx[1], trx[2]
        balanceTable[x] -= z
        balanceTable[y] += z
    }
    
    debts := make([]int, 0)
    for _, balance := range balanceTable {
        if balance != 0 {
            debts = append(debts, balance)
        }
    }
    
    return dfs(debts, 0)
}

func dfs(debts []int, idx int) int {
    for idx < len(debts) && debts[idx] == 0 {
        idx++
    }
    if idx == len(debts) {
        return 0
    }
    res := math.MaxInt32
    for i := idx+1; i < len(debts); i++ {
        if debts[idx] * debts[i] < 0 {
            debts[i] += debts[idx]
            res = min(res, 1+dfs(debts, idx+1))
            debts[i] -= debts[idx]
        }
    }
    return res
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
