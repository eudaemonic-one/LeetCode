class Solution {
    public int maxProfit(int[] prices) {
        int n = prices.length;
        if (n <= 1) {
            return 0;
        }
        int leftMin = prices[0], rightMax = prices[n-1];
        int[] left = new int[n];
        int[] right = new int[n+1];
        for (int l = 1; l < n; l++) {
            left[l] = Math.max(left[l-1], prices[l] - leftMin);
            leftMin = Math.min(leftMin, prices[l]);
            int r = n - l - 1;
            right[r] = Math.max(right[r+1], rightMax - prices[r]);
            rightMax = Math.max(rightMax, prices[r]);
        }
        int res = 0;
        for (int i = 0; i < n; i++) {
            res = Math.max(res, left[i] + right[i+1]);
        }
        return res;
    }
}