type NumMatrix struct {
    nums  [][]int
    tree    [][]int
    m       int
    n       int
}


func Constructor(matrix [][]int) NumMatrix {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return NumMatrix{}
    }
    m, n := len(matrix), len(matrix[0])
    nums := make([][]int, m)
    for i := 0; i < m; i++ {
        nums[i] = make([]int, n)
    }
    tree := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        tree[i] = make([]int, n+1)
    }
    numMatrix := NumMatrix{nums, tree, m, n}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            numMatrix.Update(i, j, matrix[i][j])
        }
    }
    return numMatrix
}


func (this *NumMatrix) Update(row int, col int, val int)  {
    delta := val - this.nums[row][col]
    this.nums[row][col] = val
    for i := row+1; i < this.m+1; i += (i & -i) {
        for j := col+1; j < this.n+1; j += (j & -j) {
            this.tree[i][j] += delta
        }
    }
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    regionSum := 0
    regionSum += getSum(this.tree, row2+1, col2+1)
    regionSum -= getSum(this.tree, row1, col2+1)
    regionSum -= getSum(this.tree, row2+1, col1)
    regionSum += getSum(this.tree, row1, col1)
    return regionSum
}

func getSum(tree [][]int, row, col int) int {
    sum := 0
    for i := row; i > 0; i -= (i & -i) {
        for j := col; j > 0; j -= (j & -j) {
            sum += tree[i][j]
        }
    }
    return sum
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * obj.Update(row,col,val);
 * param_2 := obj.SumRegion(row1,col1,row2,col2);
 */
