class Solution {
    public int maxProfit(int[] prices) {
        if (prices.length <= 1) {
            return 0;
        }
        int sold = Integer.MIN_VALUE, held = Integer.MIN_VALUE, reset = 0;
        for (int price : prices) {
            int preSold = sold;
            sold = held + price;
            held = Math.max(held, reset - price);
            reset = Math.max(reset, preSold);
        }
        return Math.max(sold, reset);
    }
}