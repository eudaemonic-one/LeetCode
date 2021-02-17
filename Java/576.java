class Solution {
    public int findPaths(int m, int n, int N, int i, int j) {
        int MOD = 1_000_000_000 + 7;
        int[][] dp = new int[m][n];
        dp[i][j] = 1;
        int count = 0;
        for (int move = 0; move < N; move++) {
            int[][] temp = new int[m][n];
            for (int r = 0; r < m; r++) {
                for (int c = 0; c < n; c++) {
                    if (r == 0) {
                        count = (count + dp[r][c]) % MOD;
                    }
                    if (r == m - 1) {
                        count = (count + dp[r][c]) % MOD;
                    }
                    if (c == 0) {
                        count = (count + dp[r][c]) % MOD;
                    }
                    if (c == n - 1) {
                        count = (count + dp[r][c]) % MOD;
                    }
                    temp[r][c] = (
                        ((r > 0 ? dp[r - 1][c] : 0) + (r < m - 1 ? dp[r + 1][c] : 0)) % MOD +
                        ((c > 0 ? dp[r][c - 1] : 0) + (c < n - 1 ? dp[r][c + 1] : 0)) % MOD
                    ) % MOD;
                }
            }
            dp = temp;
        }
        return count;
    }
}