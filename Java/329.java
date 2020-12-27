class Solution {
    private static final int[][] dirs = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};
    
    public int longestIncreasingPath(int[][] matrix) {
        int res = 0;
        int m = matrix.length;
        if (m == 0) {
            return res;
        }
        int n = matrix[0].length;
        if (n == 0) {
            return res;
        }
        int[][] memo = new int[m][n];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                res = Math.max(res, dfs(matrix, m, n, i, j, memo));
            }
        }
        return res;
    }
    
    private int dfs(int[][] matrix, int m, int n, int r, int c, int[][] memo) {
        if (memo[r][c] != 0) {
            return memo[r][c];
        }
        for (int[] dir : dirs) {
            int nr = r + dir[0];
            int nc = c + dir[1];
            if (0 <= nr && nr < m && 0 <= nc && nc < n && matrix[nr][nc] > matrix[r][c]) {
                memo[r][c] = Math.max(memo[r][c], dfs(matrix, m, n, nr, nc, memo));
            }
        }
        memo[r][c] += 1;
        return memo[r][c];
    }
}