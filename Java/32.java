class Solution {
    public int longestValidParentheses(String s) {
        if (s == null || s.length() == 0) {
            return 0;
        }
        int n = s.length();
        int[] dp = new int[n];
        int res = 0;
        for (int i = 1; i < n; i++) {
            if (s.charAt(i) == ')') {
                if (s.charAt(i-1) == '(') {
                    dp[i] = (i >= 2 ? dp[i-2] : 0) + 2;
                } else if (i-dp[i-1] > 0 && s.charAt(i-dp[i-1]-1) == '(') {
                    dp[i] = dp[i-1] + ((i-dp[i-1]) >= 2 ? dp[i-dp[i-1]-2] : 0) + 2;
                }
                res = Math.max(res, dp[i]);
            }
        }
        return res;
    }
}