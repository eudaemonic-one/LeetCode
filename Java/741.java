class Solution {
    public int cherryPickup(int[][] grid) {
        int n = grid.length;
        int[][][] memo = new int[n][n][n];
        for (int[][] layer: memo) {
            for (int[] row: layer) {
                Arrays.fill(row, Integer.MIN_VALUE);
            }
        }
        return Math.max(0, dp(grid, n, memo, 0, 0, 0));
    }
    
    private int dp(int[][] grid, int n, int[][][] memo, int r1, int c1, int c2) {
        int r2 = r1 + c1 - c2;
        if (r1 == n || r2 == n || c1 == n || c2 == n || grid[r1][c1] == -1 || grid[r2][c2] == -1) {
            return Integer.MIN_VALUE + 1;
        } else if (r1 == n-1 && c1 == n-1) {
            return grid[r1][c1];
        } else if (memo[r1][c1][c2] != Integer.MIN_VALUE) {
            return memo[r1][c1][c2];
        } else {
            int res = grid[r1][c1];
            if (c1 != c2) {
                res += grid[r2][c2];
            }
            res += Math.max(Math.max(dp(grid, n, memo, r1, c1+1, c2+1), dp(grid, n, memo, r1+1, c1, c2+1)),
                            Math.max(dp(grid, n, memo, r1, c1+1, c2), dp(grid, n, memo, r1+1, c1, c2)));
            memo[r1][c1][c2] = res;
            return res;
        }
    }
}