class Solution {
    public int maxProfit(int k, int[] prices) {
        if (k <= 0 || prices.length <= 1) {
            return 0;
        }
        int n = prices.length;
        if (2 * k > n) {
            int res = 0;
            for (int i = 1; i < n; i++) {
                if (prices[i] > prices[i-1]) {
                    res += prices[i] - prices[i-1];
                }
            }
            return res;
        }
        // dp[i][used][ishold] = balance
        int[][][] dp = new int[n][k+1][2];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j <= k; j++) {
                dp[i][j][0] = -100_000_000;
                dp[i][j][1] = -100_000_000;
            }
        }
        // initialize dp array
        dp[0][0][0] = 0;
        dp[0][1][1] = -prices[0];
        // fill the array
        for (int i = 1; i < n; i++) {
            for (int j = 0; j <= k; j++) {
                dp[i][j][0] = Math.max(dp[i-1][j][0], dp[i-1][j][1] + prices[i]);
                if (j > 0) {
                    dp[i][j][1] = Math.max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i]);
                }
            }
        }
        int res = 0;
        for (int j = 0; j <= k; j++) {
            res = Math.max(res, dp[n-1][j][0]);
        }
        return res;
    }
}