class Solution {
    public int largest1BorderedSquare(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        int[][] left = new int[m][n], top = new int[m][n];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1) {
                    left[i][j] = j > 0 ? left[i][j-1] + 1 : 1;
                    top[i][j] = i > 0 ? top[i-1][j] + 1 : 1;
                }
            }
        }
        for (int l = Math.min(m, n); l >= 1; l--) {
            for (int i = l-1; i < m; i++) {
                for (int j = l-1; j < n; j++) {
                    if (left[i][j] >= l &&
                        top[i][j] >= l &&
                        left[i-l+1][j] >= l
                        && top[i][j-l+1] >= l) {
                        return l * l;
                    }
                }
            }
        }
        return 0;
    }
}