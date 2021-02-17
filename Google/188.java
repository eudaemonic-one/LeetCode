class Solution {
    public int maxProfit(int k, int[] prices) {
        int n = prices.length;
        if (n <= 1 || k <= 0) {
            return 0;
        }
        if (2 * k > n) {
            int res = 0;
            for (int i = 1; i < n; i++) {
                if (prices[i-1] < prices[i]) {
                    res += prices[i] - prices[i-1];
                }
            }
            return res;
        }
        // dp[i][used_k][is_hold]
        int[][][] dp = new int[n][k+1][2];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j <= k; j++) {
                dp[i][j][0] = -1_000_000_000;
                dp[i][j][1] = -1_000_000_000;
            }
        }
        dp[0][0][0] = 0;
        dp[0][1][1] = -prices[0];
        for (int i = 1; i < n; i++) {
            for (int j = 0; j <= k; j++) {
                dp[i][j][0] = Math.max(dp[i-1][j][0], dp[i-1][j][1] + prices[i]);
                if (j > 0) {
                    dp[i][j][1] = Math.max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i]);
                }
            }
        }
        int maxProfit = 0;
        for (int j = 0; j <= k; j++) {
            maxProfit = Math.max(maxProfit, dp[n-1][j][0]);
        }
        return maxProfit;
    }
}