func canCompleteCircuit(gas []int, cost []int) int {
    var start, totalTank, currTank int
    for i := 0; i < len(gas); i++ {
        totalTank += gas[i] - cost[i]
        currTank += gas[i] - cost[i]
        if currTank < 0 {
            start = i + 1
            currTank = 0
        }
    }
    if totalTank < 0 {
        return -1
    }
    return start
}
