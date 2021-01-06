class Solution {
    public int numberOfWays(int numPeople) {
        long mod = (long)1E9 + 7;
        long[] dp = new long[numPeople/2+1];
        dp[0] = 1;
        for (int n = 1; n <= numPeople/2; n++) {
            for (int i = 0; i < n; i++) {
                dp[n] = (dp[n] + dp[i] * dp[n-i-1]) % mod;
            }
        }
        return (int)dp[numPeople/2];
    }
}