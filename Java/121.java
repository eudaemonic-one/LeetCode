class Solution {
    public int maxProfit(int[] prices) {
        int res = 0, min = Integer.MAX_VALUE;
        for (int price: prices) {
            if (price < min) {
                min = price;
            } else if (price - min > res) {
                res = price - min;
            }
        }
        return res;
    }
}