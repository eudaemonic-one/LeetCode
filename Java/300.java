// Approach #2: Dynamic Programming with Binary Search
class Solution {
    public int lengthOfLIS(int[] nums) {
        int[] dp = new int[nums.length];
        int len = 0;
        for (int num : nums) {
            int i = Arrays.binarySearch(dp, 0, len, num);
            if (i < 0) {
                i = -(i + 1);
            }
            dp[i] = num;
            if (i == len) {
                ++len;
            }
        }
        return len;
    }
}

// Approach #1: Dynamic Programming
// class Solution {
//     public int lengthOfLIS(int[] nums) {
//         if (nums.length == 0) {
//             return 0;
//         }
//         int[] dp = new int[nums.length];
//         dp[0] = 1;
//         int res = 1;
//         for (int i = 1; i < dp.length; i++) {
//             int tmp = 0;
//             for (int j = 0; j < i; j++) {
//                 if (nums[j] < nums[i]) {
//                     tmp = Math.max(tmp, dp[j]);
//                 }
//             }
//             dp[i] = tmp + 1;
//             res = Math.max(res, dp[i]);
//         }
//         return res;
//     }
// }