class Solution {
    public int minCost(String s, int[] cost) {
        int ans = 0;
        int sumCost = 0, maxCost = 0;
        for (int i = 0; i < s.length(); i++) {
            if (i > 0 && s.charAt(i) != s.charAt(i-1)) {
                ans += sumCost - maxCost;
                maxCost = 0;
                sumCost = 0;
            }
            sumCost += cost[i];
            maxCost = Math.max(maxCost, cost[i]);
        }
        ans += sumCost - maxCost;
        return ans;
    }
}