class Solution {
    public int maxProfit(int[] prices) {
        int n = prices.length;
        if (n <= 1) {
            return 0;
        }
        int leftMin = prices[0];
        int rightMax = prices[n-1];
        int[] leftProfits = new int[n];
        int[] rightProfits = new int[n+1];
        for (int l = 1; l < n; l++) {
            leftProfits[l] = Math.max(leftProfits[l-1], prices[l] - leftMin);
            leftMin = Math.min(leftMin, prices[l]);
            int r = n - l - 1;
            rightProfits[r] = Math.max(rightProfits[r+1], rightMax - prices[r]);
            rightMax = Math.max(rightMax, prices[r]);
        }
        int maxProfit = 0;
        for (int i = 0; i < n; i++) {
            maxProfit = Math.max(maxProfit, leftProfits[i] + rightProfits[i+1]);
        }
        return maxProfit;
    }
}