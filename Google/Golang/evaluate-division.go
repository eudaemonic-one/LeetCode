func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    operands := make(map[string]string)
    times := make(map[string]float64)
    for i := 0; i < len(equations); i++ {
        op1, op2 := equations[i][0], equations[i][1]
        root1 := find(operands, times, op1)
        root2 := find(operands, times, op2)
        operands[root1] = root2 // union
        times[root1] = times[op2] * values[i] / times[op1]
    }
    res := make([]float64, len(queries))
    for i := 0; i < len(queries); i++ {
        op1, op2 := queries[i][0], queries[i][1]
        _, ok1 := operands[op1]
        _, ok2 := operands[op2]
        if !ok1 || !ok2 || find(operands, times, op1) != find(operands, times, op2) {
            res[i] = -1.0
            continue
        }
        res[i] = times[op1] / times[op2]
    }
    return res
}

func find(operands map[string]string, times map[string]float64, target string) string {
    if _, ok := operands[target]; !ok {
        operands[target] = target
        times[target] = 1.0
        return target
    }
    if operands[target] == target {
        return target
    }
    next := operands[target]
    e := find(operands, times, next)
    operands[target] = e
    times[target] = times[target] * times[next]
    return e
}
