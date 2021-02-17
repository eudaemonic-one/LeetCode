class Solution {
    public int maxProfit(int[] prices) {
        int maxProfit = 0;
        int prev = Integer.MAX_VALUE;
        for (int price : prices) {
            if (price - prev > 0) {
                maxProfit += price - prev;
                prev = price;
            }
            if (price < prev) {
                prev = price;
            }
        }
        return maxProfit;
    }
}