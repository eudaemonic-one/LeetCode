class Solution {
    public int maxProfit(int[] prices) {
        int maxProfit = 0;
        int minPrice = Integer.MAX_VALUE;
        for (int price : prices) {
            if (price - minPrice > maxProfit) {
                maxProfit = price - minPrice;
            }
            if (price < minPrice) {
                minPrice = price;
            }
        }
        return maxProfit;
    }
}