// Approach #2: Merging

class Solution {
    public int maxProfit(int k, int[] prices) {
        int n = prices.length;
        if (n <= 1 || k <= 0) {
            return 0;
        }
        List<int[]> transactions = new ArrayList<>();
        int start = 0, end = 0;
        for (int i = 1; i < n; i++) {
            if (prices[i-1] <= prices[i]) {
                end = i;
            } else {
                if (start < end) {
                    transactions.add(new int[]{start, end});
                }
                start = i;
            }
        }
        if (start < end) {
            transactions.add(new int[]{start, end});
        }
        
        while (transactions.size() > k) {
            int deleteIndex = 0;
            int minDeleteLoss = Integer.MAX_VALUE;
            for (int i = 0; i < transactions.size(); i++) {
                int[] tx = transactions.get(i);
                int profitLoss = prices[tx[1]] - prices[tx[0]];
                if (profitLoss < minDeleteLoss) {
                    minDeleteLoss = profitLoss;
                    deleteIndex = i;
                }
            }
            int mergeIndex = 0;
            int minMergeLoss = Integer.MAX_VALUE;
            for (int i = 1; i < transactions.size(); i++) {
                int[] tx1 = transactions.get(i-1);
                int[] tx2 = transactions.get(i);
                int profitLoss = prices[tx1[1]] - prices[tx2[0]];
                if (profitLoss < minMergeLoss) {
                    minMergeLoss = profitLoss;
                    mergeIndex = i;
                }
            }
            if (minDeleteLoss <= minMergeLoss) {
                transactions.remove(deleteIndex);
            } else {
                int[] tx1 = transactions.get(mergeIndex-1);
                int[] tx2 = transactions.get(mergeIndex);
                tx1[1] = tx2[1];
                transactions.remove(mergeIndex);
            }
        }
        
        int maxProfit = 0;
        for (int[] tx: transactions) {
            maxProfit += prices[tx[1]] - prices[tx[0]];
        }
        return maxProfit;
    }
}

// Approach #1: Dynamic Programming

// class Solution {
//     public int maxProfit(int k, int[] prices) {
//         int n = prices.length;
//         if (n <= 1 || k <= 0) {
//             return 0;
//         }
//         if (2 * k > n) {
//             int res = 0;
//             for (int i = 1; i < n; i++) {
//                 if (prices[i-1] < prices[i]) {
//                     res += prices[i] - prices[i-1];
//                 }
//             }
//             return res;
//         }
//         // dp[i][used_k][is_hold]
//         int[][][] dp = new int[n][k+1][2];
//         for (int i = 0; i < n; i++) {
//             for (int j = 0; j <= k; j++) {
//                 dp[i][j][0] = -1_000_000_000;
//                 dp[i][j][1] = -1_000_000_000;
//             }
//         }
//         dp[0][0][0] = 0;
//         dp[0][1][1] = -prices[0];
//         for (int i = 1; i < n; i++) {
//             for (int j = 0; j <= k; j++) {
//                 dp[i][j][0] = Math.max(dp[i-1][j][0], dp[i-1][j][1] + prices[i]);
//                 if (j > 0) {
//                     dp[i][j][1] = Math.max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i]);
//                 }
//             }
//         }
//         int maxProfit = 0;
//         for (int j = 0; j <= k; j++) {
//             maxProfit = Math.max(maxProfit, dp[n-1][j][0]);
//         }
//         return maxProfit;
//     }
// }