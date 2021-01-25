class NumMatrix {
    private int rows;
    private int cols;
    private int bit[][];

    public NumMatrix(int[][] matrix) {
        rows = matrix.length;
        if (rows == 0) {
            return;
        }
        cols = matrix[0].length;
        bit = new int[rows+1][cols+1];
        for (int i = 1; i <= rows; ++i) {
            for (int j = 1; j <= cols; ++j) {
                int val = matrix[i - 1][j - 1];
                updateBIT(i, j, val);
            }
        }
    }
    
    private void updateBIT(int r, int c, int val) {
        for (int i = r; i <= rows; i += lsb(i)) {
            for (int j = c; j <= cols; j += lsb(j)) {
                this.bit[i][j] += val;
            }
        }
    }
    
    private int queryBIT(int r, int c) {
        int sum = 0;
        for (int i = r; i > 0; i -= lsb(i)) {
            for (int j = c; j > 0; j -= lsb(j)) {
                sum += this.bit[i][j];
            }
        }
        return sum;
    }
    
    private int lsb(int n) {
        return n & -n;
    }

    public void update(int row, int col, int val) {
        int oldVal = sumRegion(row, col, row, col);
        int diff = val - oldVal;
        updateBIT(row+1, col+1, diff);
    }
    
    public int sumRegion(int row1, int col1, int row2, int col2) {
        int a = queryBIT(row2+1, col2+1);
        int b = queryBIT(row1, col1);
        int c = queryBIT(row2+1, col1);
        int d = queryBIT(row1, col2+1);
        return (a+b) - (c+d);
    }
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * NumMatrix obj = new NumMatrix(matrix);
 * obj.update(row,col,val);
 * int param_2 = obj.sumRegion(row1,col1,row2,col2);
 */