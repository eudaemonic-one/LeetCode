class Solution {
    private static final int MOD = 1_000_000_007;
    
    public int numWays(int steps, int arrLen) {
        if (arrLen <= 1) {
            return arrLen;
        }
        int[] dp = new int[arrLen];
        dp[0] = 1;
        dp[1] = 1;
        for (int i = 1; i < steps; i++) {
            int[] tmp = new int[arrLen];
            for (int j = 0; j < Math.min(arrLen, steps-i+1); j++) {
                long ways = dp[j];
                if (j == 0) {
                    ways = (ways + dp[j+1]) % MOD;
                } else if (j == arrLen-1) {
                    ways = (ways + dp[j-1]) % MOD;
                } else {
                    ways = (dp[j-1] + ways + dp[j+1]) % MOD;
                }
                tmp[j] = (int)ways;
            }
            dp = tmp;
        }
        return (int)dp[0];
    }
}