class Solution {
    public int maxNonOverlapping(int[] nums, int target) {
        Map<Integer, Integer> dp = new HashMap<>();
        dp.put(0, 0);
        int res = 0, sum = 0;
        for (int i = 0; i < nums.length; ++i) {
            sum += nums[i];
            if (dp.containsKey(sum - target)) {
                res = Math.max(res, dp.get(sum - target) + 1);
            }
            dp.put(sum, res);
        }
        return res;
    }
}